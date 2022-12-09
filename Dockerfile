
# Docker image versions
ARG go_ver=v1.18.6-alpine3.16.2

# Docker images
ARG go_img=ghcr.io/dopos/golang-alpine

FROM ${go_img}:${go_ver}

RUN apk add --no-cache git curl

WORKDIR /build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=`git describe --tags --always`" -a ./cmd/showonce

FROM scratch

MAINTAINER Alexey Kovrizhkin <lekovr+dopos@gmail.com>

LABEL \
  org.opencontainers.image.title="showonce" \
  org.opencontainers.image.description="ShowOnce. One-time secrets service" \
  org.opencontainers.image.authors="Alexey Kovrizhkin <ak+it@elfire.ru>" \
  org.opencontainers.image.url="https://it.elfire.ru/itc/showonce" \
  org.opencontainers.image.licenses="MIT"

WORKDIR /
COPY --from=0 /build/showonce .

# Need for SSL
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080
ENTRYPOINT ["/showonce"]

