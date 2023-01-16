# 1. Stage for building the app for developement usage (Docker compose or simple docker build command)
FROM golang:1.18-alpine AS build

ARG GIT_BRANCH
ARG GIT_SHA
ARG GIT_TAG
ARG BUILD_TIMESTAMP
ARG BUILD_INFO_PKG

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /src

WORKDIR /src

COPY go.mod go.sum /src/

RUN apk add git make
RUN go mod download

COPY . /src
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "-w -s \
    -X ${BUILD_INFO_PKG}.VCSRef=${GIT_SHA} \
    -X ${BUILD_INFO_PKG}.Version=${GIT_BRANCH}:${GIT_TAG} \
    -X ${BUILD_INFO_PKG}.Date=${BUILD_TIMESTAMP}"

# 2. Stage for running the app build in stage 1 for using in developement (docker compose or simple docker build)
FROM alpine:3.11 AS dev-build

ENV TZ=Asia/Tehran \
    PATH="/app:${PATH}"

RUN apk add --update tzdata ca-certificates bash && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone && \
    mkdir -p /app && \
    chgrp -R 0 /app && \
    chmod -R g=u /app

WORKDIR /app

COPY --from=build /src/cryptos /app

CMD ["./cryptos", "serve", "-c", "config.yml"]