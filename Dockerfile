FROM golang:1.21-alpine AS builder
MAINTAINER  "Cruii <cruii811@gmail.com>"
WORKDIR /app
ENV GOPROXY https://goproxy.cn
COPY . .
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o drone-ci-feishu .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /app/drone-ci-feishu /bin

ENTRYPOINT ["/bin/drone-ci-feishu"]