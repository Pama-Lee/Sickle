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



The purpose of this project is to integrate webhooks from various platforms, such as GitHub, Slack, etc., and forward them to different integrated platforms, such as IFTTT, Discord, Wecom Bot, etc.

## Features

- Supports multiple webhook sources and target platforms
- Supports customization of webhook formats and content
- Supports filtering and transforming webhook data
- Supports logging and error handling

## Installation

You can clone this repository to your local machine using the following command:

```bash
git clone https://github.com/Pama-Lee/Sickle.git

## Install Dependencies

```bash
go mod tidy
```

## Configuration

1. Configuration using a config file

You can configure your webhook sources and target platforms in ./config/{project}.json, for example:

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

2. Configuration through a web page

## Running

```bash
go run main.go
```

## Contribution

### Contribution Template
```markdown
# Contributor
name: Pama-Lee

# Contribution Details
- [x] Feature 1
- [ ] Feature 2
```





