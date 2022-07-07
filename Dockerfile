
FROM golang:1.18.3

#FROM ghcr.io/dopos/golang-alpine:v1.16.15-alpine3.14.3
#RUN apk add --no-cache git curl

ENV APP_VERSION 0.1.0

WORKDIR /opt/showonce
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
COPY --from=0 /opt/showonce/showonce .
# Need for SSL
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080
ENTRYPOINT ["/showonce"]

