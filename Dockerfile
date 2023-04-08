# syntax=docker/dockerfile:1

## Build
FROM golang:1.19.7-bullseye AS build

ENV GOPROXY https://goproxy.cn

WORKDIR /restful-provider

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY pkg ./pkg
COPY mock ./mock

RUN go build -o ./mock-server ./mock/main.go

## Deploy
FROM debian:stable-slim
ENV TZ 'Asia/Shanghai'
WORKDIR /

COPY --from=build /restful-provider/mock-server /restful-provider/mock/*.yaml ./
EXPOSE 8080
CMD ["/bin/sh", "-c", "exec /mock-server"]