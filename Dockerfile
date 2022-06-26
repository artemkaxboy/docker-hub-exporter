FROM golang:1.18-alpine as builder

ADD app /build/app
WORKDIR /build/app

ENV GOFLAGS="-mod=vendor"
ENV CGO_ENABLED=0

RUN echo go version: `go version`

RUN \
    go test ./... && \
    go build -o docker-hub-exporter /build/app

FROM alpine:3.16.0

ENV \
    TZ=Europe/London  \
    APP_USER=appuser  \
    APP_UID=1000

RUN \
    apk add --no-cache --update tzdata curl ca-certificates && \
    mkdir -p /usr/local/sbin && ln -s /usr/sbin/addgroup /usr/local/sbin/ && \
    adduser -s /bin/sh -D -u $APP_UID $APP_USER && chown -R $APP_USER:$APP_USER /home/$APP_USER && \
    rm -rf /var/cache/apk/*

ARG VERSION=SNAPSHOT
ARG REVISION=LOCAL
ARG REF_NAME
ARG CREATED

# https://github.com/opencontainers/image-spec/blob/main/annotations.md
LABEL org.opencontainers.image.created=${CREATED}
LABEL org.opencontainers.image.authors="artemkaxboy@gmail.com"
LABEL org.opencontainers.image.url="https://github.com/artemkaxboy/docker-hub-exporter"
LABEL org.opencontainers.image.documentation="https://github.com/artemkaxboy/docker-hub-exporter"
LABEL org.opencontainers.image.source="https://github.com/artemkaxboy/docker-hub-exporter"
LABEL org.opencontainers.image.version=${VERSION}
LABEL org.opencontainers.image.revision=${REVISION}
LABEL org.opencontainers.image.vendor="artemkaxboy@gmail.com"
LABEL org.opencontainers.image.licenses="MIT License"
LABEL org.opencontainers.image.ref.name=${REF_NAME}
LABEL org.opencontainers.image.title="Docker Hub Exporter"
LABEL org.opencontainers.image.description="Prometheus exporter for Docker Hub"


COPY --from=builder /build/app/docker-hub-exporter /srv/docker-hub-exporter
RUN chown -R $APP_USER:$APP_USER /srv
WORKDIR /srv

USER $APP_USER:$APP_USER

EXPOSE 9170
HEALTHCHECK --interval=30s --timeout=3s CMD curl --fail http://localhost:9170/ping || exit 1

ENTRYPOINT ["/srv/docker-hub-exporter", "server"]
