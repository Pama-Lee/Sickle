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

    # 数据持久化
    volume "sickle-data" {
      type      = "host"
      source    = "sickle-data"
      read_only = false
    }

    task "sickle" {
      driver = "docker"

      config {
        image = "pamalee2023/sickle:latest"
        ports = ["http"]

        # 挂载配置和数据目录
        volumes = [
          "local/config:/app/config",
          "local/webhooks:/app/webhooks",
        ]
      }

      # 挂载持久化卷
      volume_mount {
        volume      = "sickle-data"
        destination = "/app/logs"
        read_only   = false
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

      # Webhook 配置模板（示例）
      template {
        data = <<EOF
{
    "name": "Example Webhook",
    "url": "",
    "uuid": "",
    "source": {
        "type": "github",
        "config": {
            "events": [
                {
                    "name": "push",
                    "key": "sickle.headers.X-GitHub-Event",
                    "destinations": ["notification"]
                }
            ]
        }
    },
    "destinations": [
        {
            "name": "notification",
            "type": "webhook",
            "config": {
                "webhook_url": "https://your-webhook-url.com",
                "method": "POST",
                "headers": {
                    "Content-Type": "application/json"
                },
                "data": [
                    {
                        "event": "push",
                        "data": {
                            "text": "New push from ${repository.full_name}: ${head_commit.message}"
                        }
                    }
                ]
            }
        }
    ]
}
EOF

        destination = "local/webhooks/example.json"
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
