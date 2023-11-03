## template Makefile:
## service example
#:

SHELL          = /bin/sh
CFG           ?= .env
PRG            = showonce
PRG           ?= $(shell basename $$PWD)
PRG_DEST      ?= $(PRG)

# -----------------------------------------------------------------------------
# Build config

GO            ?= go
GOLANG_VERSION = v1.19.7-alpine3.17.2

SOURCES        = $(shell find . -maxdepth 3 -mindepth 1 -path ./var -prune -o -name '*.go')
APP_VERSION   ?= $(shell git describe --tags --always)
# Last project tag (used in `make changelog`)
RELEASE       ?= $(shell git describe --tags --abbrev=0 --always)
# Repository address (compiled into main.repo)
REPO          ?= $(shell git config --get remote.origin.url)

TARGETOS      ?= linux
TARGETARCH    ?= amd64
LDFLAGS       := -s -w -extldflags '-static'

ALLARCH       ?= "linux/amd64 linux/386 darwin/amd64 linux/arm linux/arm64"
DIRDIST       ?= dist

# Path to golang package docs
GODOC_REPO    ?= github.com/!le!kovr/$(PRG)
# App docker image
DOCKER_IMAGE  ?= ghcr.io/lekovr/$(PRG)

# -----------------------------------------------------------------------------
# Docker image config

#- App name
APP_NAME      ?= $(PRG)

#- Docker-compose project name (container name prefix)
PROJECT_NAME  ?= $(PRG)

# Hardcoded in docker-compose.yml service name
DC_SERVICE    ?= app

#- Docker image name
IMAGE         ?= $(DOCKER_IMAGE)

#- Docker image tag
IMAGE_VER     ?= latest

# -----------------------------------------------------------------------------
# App config

#- Docker container addr:port
LISTEN        ?= :8080

#- Public GRPC service addr:port
LISTEN_GRPC   ?= :8081

# app url prefix
APP_PROTO     ?= http

#- Auth service type
AS_TYPE       ?= gitea
#- Auth service URL
AS_HOST       ?= https://git.vivo.sb
#- Auth service org
AS_TEAM       ?= dcape
#- Auth service client_id
AS_CLIENT_ID  ?= you_should_get_id_from_as
#- Auth service client key
AS_CLIENT_KEY ?= you_should_get_key_from_as

#- Auth service cookie sign key
AS_COOKIE_SIGN_KEY   ?= $(shell < /dev/urandom tr -dc A-Za-z0-9 | head -c32; echo)
#- Auth service cookie crypt key
AS_COOKIE_CRYPT_KEY  ?= $(shell < /dev/urandom tr -dc A-Za-z0-9 | head -c32; echo)

# ------------------------------------------------------------------------------
-include $(CFG).bak
-include $(CFG)
export

# Find and include DCAPE/apps/drone/dcape-app/Makefile
DCAPE_COMPOSE ?= dcape-compose
DCAPE_ROOT    ?= $(shell docker inspect -f "{{.Config.Labels.dcape_root}}" $(DCAPE_COMPOSE))

ifeq ($(shell test -e $(DCAPE_ROOT)/Makefile.app && echo -n yes),yes)
  include $(DCAPE_ROOT)/Makefile.app
endif


.PHONY: build build-standalone run fmt lint ci-lint vet test cov-html cov-func cov-total cov-clean changelog
.PHONY: buildall dist clean docker docker-multi use-own-hub godoc ghcr

# ------------------------------------------------------------------------------
## Compile operations
#:

## Build app
build: $(PRG_DEST)

$(PRG_DEST): $(SOURCES) go.mod go.sum
	GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
	  $(GO) build -v -o $@ -ldflags \
	 "${LDFLAGS}-X main.version=$(APP_VERSION) -X main.repo=$(REPO)" \
	 ./cmd/$(PRG)

## Build like docker image from scratch
build-standalone: CGO_ENABLED=0
build-standalone: test
build-standalone: $(PRG_DEST)

## Build & run app
run: $(PRG)
	./$(PRG) --log.debug --root=static --listen $(LISTEN) --listen_grpc $(LISTEN_GRPC)

## Format go sources
fmt:
	$(GO) fmt ./...

## Check and sort imports
fmt-gci:
	@which gci > /dev/null || $(GO) install github.com/daixiang0/gci@latest
	@gci write . --skip-generated -s standard -s default

## Run lint
lint:
	@which golint > /dev/null || $(GO) install golang.org/x/lint/golint@latest
	@golint ./...

## Run golangci-lint
ci-lint:
	@golangci-lint run ./...

## Run vet
vet:
	@$(GO) vet ./...

## Run tests
test: lint vet coverage.out

test-race:
	$(GO) test -tags test -race -covermode=atomic -coverprofile=$@ ./...

# internal target
coverage.out: $(SOURCES)
	@#GIN_MODE=release $(GO) test -test.v -test.race -coverprofile=$@ -covermode=atomic ./...
	$(GO) test -tags test -covermode=atomic -coverprofile=$@ ./...

## Open coverage report in browser
cov-html: cov
	$(GO) tool cover -html=coverage.out

## Show code coverage per func
cov-func: coverage.out
	$(GO) tool cover -func coverage.out

## Show total code coverage
cov-total: coverage.out
	@$(GO) tool cover -func coverage.out | grep total: | awk '{print $$3}'

## Clean coverage report
cov-clean:
	rm -f coverage.*

## count LoC without generated code
cloc:
	@cloc --md --fullpath --exclude-dir=zgen --not-match-f=./proto/README.md \
	  --not-match-f=static/js/api.js --not-match-f=static/js/service.swagger.json  .

## Changes from last tag
changelog:
	@echo Changes since $(RELEASE)
	@echo
	@git log $(RELEASE)..@ --pretty=format:"* %s"

# ------------------------------------------------------------------------------
## GRPC operations
#:

BUF_IMG ?= ghcr.io/apisite/gogens
#BUF_IMG ?= buf

## Generate files
buf-gen:
	docker run --rm  -v `pwd`:/mnt/pwd -w /mnt/pwd $(BUF_IMG) --debug generate --template buf.gen.yaml --path proto

## Run buf command
buf-cmd:
	docker run --rm -it  -v `pwd`:/mnt/pwd -w /mnt/pwd $(BUF_IMG) $(CMD)

## Run sh command
buf-sh:
	docker run --rm -it --entrypoint /bin/sh -v `pwd`:/mnt/pwd -w /mnt/pwd $(BUF_IMG) $(CMD)

.PHONY: buf.lock

## Fetch buf.lock
buf.lock:
	@id=$$(docker create $(BUF_IMG)) ; \
	docker cp $$id:/app/$@ $@ ; \
	docker rm -v $$id

## Generate JS API
js: static/js/api.js

static/js/api.js: zgen/ts/proto/service.pb.ts
	docker run --rm  -v `pwd`:/mnt/pwd -w /mnt/pwd --entrypoint /go/bin/esbuild $(BUF_IMG)  \
	  zgen/ts/proto/service.pb.ts --bundle --outfile=/mnt/pwd/static/js/api.js --global-name=AppAPI

# ------------------------------------------------------------------------------
# GRPC testing


GRPC_HOST ?= grpc.showonce.dev.test:443

grpc-test-list:
	docker run --rm -v /etc/ssl/certs:/etc/ssl/certs -t fullstorydev/grpcurl $(GRPC_HOST) list

grpc-test-desc:
	docker run --rm -v /etc/ssl/certs:/etc/ssl/certs -t fullstorydev/grpcurl \
	 $(GRPC_HOST) describe api.showonce.v1.PublicService

ID ?= 01H6DWS3VFV0YXEG0B25BSQ0R9

grpc-test-data:
	docker run --rm -v /etc/ssl/certs:/etc/ssl/certs -t fullstorydev/grpcurl \
	  -d '{"id": "$(ID)"}'  \
	  $(GRPC_HOST) api.showonce.v1.PublicService.GetMetadata

# ------------------------------------------------------------------------------
## Prepare distros
#:

## build app for all platforms
buildall: lint vet
	@echo "*** $@ ***" ; \
	  for a in "$(ALLARCH)" ; do \
	    echo "** $${a%/*} $${a#*/}" ; \
	    P=$(PRG)-$${a%/*}_$${a#*/} ; \
	    $(MAKE) -s build-standalone TARGETOS=$${a%/*} TARGETARCH=$${a#*/} PRG_DEST=$$P ; \
	  done

## create disro files
dist: clean buildall
	@echo "*** $@ ***"
	@[ -d $(DIRDIST) ] || mkdir $(DIRDIST)
	@sha256sum $(PRG)-* > $(DIRDIST)/SHA256SUMS ; \
	  for a in "$(ALLARCH)" ; do \
	    echo "** $${a%/*} $${a#*/}" ; \
	    P=$(PRG)-$${a%/*}_$${a#*/} ; \
	    zip "$(DIRDIST)/$$P.zip" "$$P" README.md README.ru.md screenshot.png; \
	    rm "$$P" ; \
	  done

## clean generated files
clean:
	@echo "*** $@ ***" ; \
	  for a in "$(ALLARCH)" ; do \
	    P=$(PRG)_$${a%/*}_$${a#*/} ; \
	    [ -f $$P ] && rm $$P || true ; \
	  done
	@[ -d $(DIRDIST) ] && rm -rf $(DIRDIST) || true
	@[ -f $(PRG) ] && rm -f $(PRG) || true
	@rm -f $(PRG)-* || true
	@[ ! -f coverage.out ] || rm coverage.out

# ------------------------------------------------------------------------------
## Docker build operations
#:

# build docker image directly
docker: $(PRG)
	docker build -t $(PRG) .

ALLARCH_DOCKER ?= "linux/amd64,linux/arm/v7,linux/arm64"

# build multiarch docker images via buildx
docker-multi:
	time docker buildx build --platform $(ALLARCH_DOCKER) -t $(DOCKER_IMAGE):$(APP_VERSION) --push .

OWN_HUB ?= it.elfire.ru

buildkit.toml:
	@echo [registry."$(OWN_HUB)"] > $@
	@echo ca=["/etc/docker/certs.d/$(OWN_HUB)/ca.crt"] >> $@

use-own-hub: buildkit.toml
	@docker buildx create --use --config $<

# ------------------------------------------------------------------------------
## Other
#:

## update docs at pkg.go.dev
godoc:
	vf=$(APP_VERSION) ; v=$${vf%%-*} ; echo "Update for $$v..." ; \
	curl 'https://proxy.golang.org/$(GODOC_REPO)/@v/'$$v'.info'

## update latest docker image tag at ghcr.io
ghcr:
	v=$(APP_VERSION) ; echo "Update for $$v..." ; \
	docker pull $(DOCKER_IMAGE):$$v && \
	docker tag $(DOCKER_IMAGE):$$v $(DOCKER_IMAGE):latest && \
	docker push $(DOCKER_IMAGE):latest

# ------------------------------------------------------------------------------

# Load AUTH_TOKEN
-include $(DCAPE_ROOT)/var/oauth2-token

# create OAuth application credentials
oauth2-create:
	$(MAKE) -s oauth2-app-create HOST=$(AS_HOST) URL=/login PREFIX=AS
