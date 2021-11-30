FROM golang:1.16-alpine AS builder
LABEL MAINTAINER="minhtran.in@outlook.com"

RUN mkdir -p /go/src/user_svc
WORKDIR /go/src/user_svc

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# RUN go get github.com/canthefason/go-watcher
# RUN go get github.com/pilu/fresh
# RUN go install github.com/canthefason/go-watcher/cmd/watcher
EXPOSE 8081
ENTRYPOINT  fresh
