```
__       _____ _      __   __         __
\ \     / ___/(_)____/ /__/ /__     _/ /
 \ \    \__ \/ / ___/ //_/ / _ \   / __/
 / /   ___/ / / /__/ ,< / /  __/  (_  )
/_/   /____/_/\___/_/|_/_/\___/  /  _/
                                 /_/
```

![visitors](https://badges.strrl.dev/visits/Pama-Lee/Sickle)

![chart](./image/chart.svg)

本项目的作用是统合各平台的webhooks，例如GitHub, Slack等，并且转发到不同平台的集成平台，例如IFTTT, Discord, Wecom Bot等

## 功能

- 支持多种webhooks来源和目标平台
- 支持自定义webhooks的格式和内容
- 支持过滤和转换webhooks的数据
- 支持日志和错误处理

## 安装

你可以使用以下命令克隆这个仓库到你的本地：

```bash
git clone https://github.com/Pama-Lee/Sickle.git
```

## 安装依赖

```bash
go mod tidy
```

## 配置

1. 使用config文件配置

你可以在`./config/{project}.json`中配置你的webhooks来源和目标平台，例如：

```json
// ./config/webhooks.json
{
  "webhooks": [
    {
      "name": "Webhook 1",
      // Generated after running
      "url": "",
      "source": {
        "type": "source_platform",
        "config": {
          "events": [
            {
              "name": "event1",
              "key": "data.event.event_name",
              "destinations": ["destination_platform_1"]
            },
            {
              "name": "event2",
              "key": "data.event.event_name",
              "destinations": ["destination_platform_2", "destination_platform_3"]
            }
          ]
        }
      },
      "destinations": [
        {
          "name": "destination_platform_1",
          "type": "destination_platform_1",
          "config": {
            "webhook_url": "https://destination1.com",
            "data": {
              "custom": "${data.event.data.value1}",
              "custom2": "${data.event.data.value2}"
            }
          }
        },
        {
          "name": "destination_platform_2",
          "type": "destination_platform_2",
          "config": {
            "webhook_url": "https://destination2.com",
            "data": {
              "custom": "${data.event.data.value1}",
              "custom2": "${data.event.data.value2}"
            }
          }
        },
        {
          "name": "destination_platform_3",
          "type": "destination_platform_3",
          "config": {
            "webhook_url": "https://destination3.com",
            "data": {
              "custom": "${data.event.data.value1}",
              "custom2": "${data.event.data.value2}"
            }
          }
        }
      ]
    }
  ]
}

```

2. 在web页面配置

## 运行

```bash
go run main.go
```

## 贡献

### 贡献模板
```markdown
# 贡献者
name: Pama-Lee

# 贡献内容
- [x] 功能1
- [ ] 功能2
```
