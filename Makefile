## template Makefile:
## service example
#:

SHELL          := /bin/sh
CFG            := .env
CFG_TMPL       := Makefile.env
PRG            := $(shell basename $$PWD)

# -----------------------------------------------------------------------------
# Docker image config

#- App name
APP_NAME      ?= $(PRG)

#- Docker-compose project name (container name prefix)
PROJECT_NAME  ?= $(PRG)

# Hardcoded in docker-compose.yml service name
DC_SERVICE    ?= app

#- Docker image name
IMAGE         := $(if $(IMAGE),$(IMAGE),$(DOCKER_IMAGE))

#- Docker image tag
IMAGE_VER     ?= latest

# -----------------------------------------------------------------------------
# App config

-include $(CFG_TMPL)

AS_CLIENT_ID := $(if $(AS_CLIENT_ID),$(AS_CLIENT_ID),you_should_get_id_from_as)
AS_CLIENT_KEY := $(if $(AS_CLIENT_KEY),$(AS_CLIENT_KEY),you_should_get_id_from_as)

AS_COOKIE_SIGN_KEY  := $(if $(AS_COOKIE_SIGN_KEY),$(AS_COOKIE_SIGN_KEY),$(shell < /dev/urandom tr -dc A-Za-z0-9 | head -c32; echo))
AS_COOKIE_CRYPT_KEY := $(if $(AS_COOKIE_CRYPT_KEY),$(AS_COOKIE_CRYPT_KEY),$(shell < /dev/urandom tr -dc A-Za-z0-9 | head -c32; echo))

# used in URL generation
ifeq ($(USE_TLS),false)
#- app url prefix
HTTP_PROTO := http
else
HTTP_PROTO := https
endif

# ------------------------------------------------------------------------------
-include $(CFG).bak
-include $(CFG)
export

include Makefile.golang

# Find and include DCAPE/apps/drone/dcape-app/Makefile
DCAPE_COMPOSE := dcape-compose
DCAPE_ROOT    := $(shell docker inspect -f "{{.Config.Labels.dcape_root}}" $(DCAPE_COMPOSE))

ifeq ($(shell test -e $(DCAPE_ROOT)/Makefile.app && echo -n yes),yes)
  include $(DCAPE_ROOT)/Makefile.app
endif

.PHONY: buildall dist clean docker docker-multi use-own-hub godoc ghcr

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
