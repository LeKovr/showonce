ARG GOLANG_IMAGE=ghcr.io/dopos/golang-alpine
ARG GOLANG_VERSION=v1.22.3-alpine3.20
ARG APP=showonce

FROM --platform=$BUILDPLATFORM ${GOLANG_IMAGE}:${GOLANG_VERSION} AS build

ARG APP
ARG GOPROXY TARGETOS TARGETARCH

RUN apk add --no-cache curl git make jq bash openssl

COPY . /src/$APP
WORKDIR /src/$APP

RUN --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
    --mount=type=cache,id=gomod,target=/go/pkg \
    make -f Makefile.golang build PRG_DEST=$APP-static

FROM scratch

ARG APP

LABEL org.opencontainers.image.title="$APP" \
      org.opencontainers.image.description="ShowOnce. One-time secrets service" \
      org.opencontainers.image.authors="lekovr+github@gmail.com" \
      org.opencontainers.image.licenses="Apache v2"

WORKDIR /

COPY --from=build /src/$APP/$APP-static /app
EXPOSE 8080
ENTRYPOINT [ "/app" ]
