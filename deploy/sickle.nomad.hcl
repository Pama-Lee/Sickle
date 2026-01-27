job "sickle" {
  datacenters = ["dc1"]
  type        = "service"

  meta {
    version = "1.0.0"
  }

  group "sickle" {
    count = 1

    network {
      port "http" {
        static = 8080
        to     = 8080
      }
    }

    task "sickle" {
      driver = "docker"

      config {
        image = "ghcr.io/pama-lee/sickle:latest"
        ports = ["http"]

        # 直接映射配置目录
        volumes = [
          "local/config:/app/config",
          "local/webhooks:/app/webhooks",
          "local/logs:/app/logs",
          "local/request:/app/request",
        ]
      }

      # 环境变量
      env {
        TZ = "Asia/Shanghai"
      }

      # 配置文件模板
      template {
        data = <<EOF
{
    "database": {
        "type": "sqlite",
        "host": "127.0.0.1",
        "port": 3306,
        "user": "",
        "password": "",
        "database": ""
    },
    "server": {
        "host": "0.0.0.0",
        "port": 8080
    },
    "log": {
        "file": "sickle",
        "level": "info"
    },
    "project": {
        "name": "Sickle",
        "version": "1.0.0",
        "author": "Pama Lee",
        "description": "A centralized system for managing multi-platform Hooks."
    }
}
EOF

        destination = "local/config/config.json"
        change_mode = "restart"
      }

      # GitHub -> 企业微信 Webhook 配置
      template {
        data = <<EOF
{
    "name": "GitHub to WeChat Work",
    "url": "",
    "uuid": "",
    "source": {
        "type": "github",
        "config": {
            "events": [
                {
                    "name": "push",
                    "key": "sickle.headers.X-Github-Event",
                    "destinations": ["wechat"]
                },
                {
                    "name": "pull_request",
                    "key": "sickle.headers.X-Github-Event",
                    "destinations": ["wechat"]
                },
                {
                    "name": "issues",
                    "key": "sickle.headers.X-Github-Event",
                    "destinations": ["wechat"]
                },
                {
                    "name": "star",
                    "key": "sickle.headers.X-Github-Event",
                    "destinations": ["wechat"]
                }
            ]
        }
    },
    "destinations": [
        {
            "name": "wechat",
            "type": "wechat_work",
            "config": {
                "webhook_url": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=fa29effe-773e-4f8f-918d-7f4663d6b480",
                "method": "POST",
                "headers": {
                    "Content-Type": "application/json"
                },
                "data": [
                    {
                        "event": "push",
                        "data": {
                            "msgtype": "markdown",
                            "markdown": {
                                "content": "## 🚀 代码推送\n> 仓库: **${repository.full_name}**\n> 分支: `${ref}`\n> 提交者: ${pusher.name}\n> 提交信息: ${head_commit.message}\n> [查看详情](${head_commit.url})"
                            }
                        }
                    },
                    {
                        "event": "pull_request",
                        "data": {
                            "msgtype": "markdown",
                            "markdown": {
                                "content": "## 🔀 Pull Request ${action}\n> 仓库: **${repository.full_name}**\n> 标题: ${pull_request.title}\n> 提交者: ${pull_request.user.login}\n> [查看详情](${pull_request.html_url})"
                            }
                        }
                    },
                    {
                        "event": "issues",
                        "data": {
                            "msgtype": "markdown",
                            "markdown": {
                                "content": "## 📝 Issue ${action}\n> 仓库: **${repository.full_name}**\n> 标题: ${issue.title}\n> 提交者: ${issue.user.login}\n> [查看详情](${issue.html_url})"
                            }
                        }
                    },
                    {
                        "event": "star",
                        "data": {
                            "msgtype": "markdown",
                            "markdown": {
                                "content": "## ⭐ Star ${action}\n> 仓库: **${repository.full_name}**\n> 用户: ${sender.login}\n> 当前 Stars: ${repository.stargazers_count}"
                            }
                        }
                    }
                ]
            }
        }
    ]
}
EOF

        destination = "local/webhooks/github-wechat.json"
        change_mode = "restart"
      }

      # 资源限制
      resources {
        cpu    = 256
        memory = 256
      }

      # 服务注册
      service {
        name = "sickle"
        port = "http"
        tags = ["webhook", "sickle"]

        check {
          name     = "sickle-health"
          type     = "tcp"
          port     = "http"
          interval = "30s"
          timeout  = "5s"
        }
      }
    }
  }
}
