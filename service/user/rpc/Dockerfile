ROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY /service/user/rpc/etc /app/etc
RUN go run service/user/rpc/user.go