FROM golang:1.15-alpine as base

RUN apk --update add build-base

WORKDIR /api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
