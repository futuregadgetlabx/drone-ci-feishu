## drone-feishu
推送DroneCI的流水线构建信息到飞书

## 使用示例

```yaml
kind: pipeline
type: docker
name: ci
steps:
  - name: build
    image: golang:1.21-alpine
    commands:
      - GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o drone-feishu .
  - name: notify
    image: cruii/drone-feishu
    pull: always
    settings:
      user_id:
        # 飞书用户ID
        from_secret: user_id
      app_id:
        #飞书机器人App ID
        from_secret: app_id
      app_secret:
        #飞书机器人App Secret
        from_secret: app_secret
trigger:
  event:
    include:
      - push
      - pull_request
  branch:
    - main
    - dev
```
