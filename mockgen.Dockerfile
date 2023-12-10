FROM golang:alpine as builder

RUN apk add --no-cache git

WORKDIR /src

COPY . .

RUN go mod download
RUN go install github.com/golang/mock/mockgen@v1.7.0-rc.1