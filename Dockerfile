FROM golang:1.21-alpine AS builder
LABEL maintainer="cruii <cruii811@gmail.com>" \
  org.label-schema.name="drone-feishu" \
  org.label-schema.vendor="cruii" \
  org.label-schema.schema-version="1.0"

LABEL org.opencontainers.image.source=https://github.com/futuregadgetlabx/drone-feishu
LABEL org.opencontainers.image.description="drone-feishu"
LABEL org.opencontainers.image.licenses=MIT
WORKDIR /app
ENV GOPROXY https://goproxy.cn
COPY . .
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o drone-feishu .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
COPY --from=builder /app/drone-feishu /bin
ENTRYPOINT ["/bin/drone-feishu"]
