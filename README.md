```
__       _____ _      __   __         __
\ \     / ___/(_)____/ /__/ /__     _/ /
 \ \    \__ \/ / ___/ //_/ / _ \   / __/
 / /   ___/ / / /__/ ,< / /  __/  (_  )
/_/   /____/_/\___/_/|_/_/\___/  /  _/
                                 /_/
```

![visitors](https://badges.strrl.dev/visits/Pama-Lee/Sickle)

本项目的作用是统合各平台的webhooks，例如GitHub, Slack, Discord等，并且转发到不同平台的集成平台，例如IFTTT, Zapier等。

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

你可以在`config.json`中配置你的webhooks来源和目标平台，例如：

```json
{
    "port": 8080,
    "webhooks": [
        {
            "name": "GitHub",
            "path": "/github",
            "method": "POST",
            "headers": {
                "Content-Type": "application/json"
            },
            "body": {
                "type": "json"
            },
            "targets": [
                {
                    "name": "IFTTT",
                    "url": "https://maker.ifttt.com/trigger/{event}/with/key/{key}",
                    "method": "POST",
                    "headers": {
                        "Content-Type": "application/json"
                    },
                    "body": {
                        "type": "json",
                        "data": {
                            "value1": "{value1}",
                            "value2": "{value2}",
                            "value3": "{value3}"
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





