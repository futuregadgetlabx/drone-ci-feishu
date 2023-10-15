FROM golang:1.21-alpine AS builder
MAINTAINER  "Cruii <cruii811@gmail.com>"
WORKDIR /app
ENV GOPROXY https://goproxy.cn
COPY . .
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o drone-feishu .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/drone-feishu .
COPY --from=builder /app/template ./template
ENTRYPOINT ["/app/drone-feishu"]
