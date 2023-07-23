ARG GOLANG_IMAGE=ghcr.io/dopos/golang-alpine
ARG GOLANG_VERSION=v1.19.7-alpine3.17.2
ARG APP=showonce

FROM --platform=$BUILDPLATFORM ${GOLANG_IMAGE}:${GOLANG_VERSION} as build

ARG APP

COPY . /src/$APP
WORKDIR /src/$APP

ARG GOPROXY TARGETOS TARGETARCH
RUN --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
    --mount=type=cache,id=gomod,target=/go/pkg \
    make build-standalone

FROM scratch

ARG APP

LABEL org.opencontainers.image.title="$APP" \
      org.opencontainers.image.description="ShowOnce. One-time secrets service" \
      org.opencontainers.image.authors="lekovr+dopos@gmail.com" \
      org.opencontainers.image.licenses="MIT"

WORKDIR /

COPY --from=build /src/$APP/$APP /app
EXPOSE 8080
ENTRYPOINT [ "/app" ]
