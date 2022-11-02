# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /usr/src/app/

RUN go build -o /99minuts cmd/main.go
