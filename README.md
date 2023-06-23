```
__       _____ _      __   __         __
\ \     / ___/(_)____/ /__/ /__     _/ /
 \ \    \__ \/ / ___/ //_/ / _ \   / __/
 / /   ___/ / / /__/ ,< / /  __/  (_  )
/_/   /____/_/\___/_/|_/_/\___/  /  _/
                                 /_/
```
[中文](./README_CN.md)

![visitors](https://badges.strrl.dev/visits/Pama-Lee/Sickle)
[![Discord](https://img.shields.io/discord/1106894340548210720?color=7289da&logo=discord&logoColor=white&style=flat-square)](https://discord.com/invite/zY2c9B78)
![chart](./image/chart.svg)

The role of this project is to integrate the webhooks of various platforms, such as GitHub, Slack, etc., and forward
them to the integration platforms of different platforms, such as IFTTT, Discord, Wecom Bot, etc.

## <img src="https://media.tenor.com/SNL9_xhZl9oAAAAj/waving-hand-joypixels.gif" width="25px"> Welcome

Welcome to Sickle, a project for integrating webhooks, this project is just started, developed in Go language, you can
find some useful information here. If you have any questions,
Please feel free to raise an issue. Or contact me by email: [Pama Lee](mailto:pama@pamalee.cn)

## Features

- Support multiple webhooks sources and target platforms
- Support for customizing the format and content of webhooks
- Support for filtering and transforming data from webhooks
- Support logging and error handling

## Install

You can clone this repository to your local with the following command:

```bash
git clone https://github.com/Pama-Lee/Sickle.git
```

## Install dependencies

```bash
go mod tidy
```

## run

1. Run from source

```bash
go run Sickle
```

2. Run from the binary
   You can compile binaries yourself, or download compiled binaries
   from [Release](https://github.com/Pama-Lee/Sickle/releases).

```bash
./Sickle
```

## 🤔How to use

1. Run Sickle
   It is recommended to run Sickle on `Linux` or `MacOS`, you can use the `nohup` command to run Sickle in the
   background:

```bash
nohup ./Sickle &
```

![nohup](./image/run1.png)

When running Sickle for the first time, a `./webhooks` folder will be created automatically to store webhooks
configuration files, where you can configure your webhooks source and target platform.
And an `example.json` file will be created in the `./webhooks` folder for configuration examples, you can refer to this
file to configure your webhooks.

2. Configure webhooks and get UUID
   After running Sickle, Sickle will automatically read the configuration files in the `./webhooks` folder, and assign a
   UUID to each configuration file, you can view your UUID in the configuration file.

A UUID is a string like this: `784db24b-e9fb-4e4f-8b7d-fac84500a9af`

3. Configure webhooks on the webhooks source platform

You can configure webhooks on the webhooks source platform, and set the target address of webhooks
to `http://your_ip:port/webhooks/webhooks/{UUID}`
> If you configured Sickle with a domain name, replace `your_ip` with your domain name

For example, if you want to configure Github's webhooks, you can configure webhooks on Github's webhooks configuration
page, and set the target address of webhooks to `http://your_ip:port/webhooks/webhooks/{UUID}`

**❗Note**: Please set `Content type` to `application/json`

![Github](./image/github.png)

4. One step closer

At this point Sickle can already receive webhooks, but Sickle doesn't know where you want to forward webhooks, you need
to configure your webhooks target platform in the configuration file in the `./webhooks` folder.

For example, if you want to forward Github's webhooks to IFTTT, you need to configure your IFTTT's webhooks in the
configuration file in the `./webhooks` folder,

For example the webhooks link for `IFTTT` is `https://maker.ifttt.com/trigger/{event}/with/key/{key}`
> Please replace `{event}` with your event name and `{key}` with your key

You can configure your IFTTT webhooks in the configuration file in the `./webhooks` folder, and set the `url`
of `webhooks` to `https://maker.ifttt.com/trigger/{event}/ with/key/{key}`

### 🎇 You're done

## configuration

1. Use the config file to configure

You can configure your webhooks source and target platform in `./webhooks/{project}.json`, for example:

```json
{
  "webhooks": [
    {
      "name": "Webhook 1",
      "url": "",
      "uuid": "",
      "source": {
        "type": "source_platform",
        "config": {
          "events": [
            {
              "name": "event1",
              "key": "data. event. event_name",
              "destinations": [
                "destination_platform_1"
              ]
            },
            {
              "name": "event2",
              "key": "data. event. event_name",
              "destinations": [
                "destination_platform_1"
              ]
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
            "method": "POST",
            "headers": {
              "User-Agent": "Sickle"
            },
            "data": [
              {
                "event": "event1",
                "data": {
                  "value1": "sickle.data.event.data.value1",
                  "value2": "sickle.data.event.data.value2"
                }
              },
              {
                "event": "event2",
                "data": {
                  "value1": "sickle.data.event.data.value1",
                  "value2": "sickle.data.event.data.value2"
                }
              }
            ]
          }
        }
      ]
    }
  ]
}

```

### Configuration instructions

<mark>Is planning to make an online configuration generator</mark>

> Description of the data keys included below: You can view the data keys of your webhooks
> in `./request/{name}_{UUID}.json`. This file is automatically generated from the request body when the first webhooks
> arrive,
> Helps you locate data. If your webhooks have not arrived, you can use the following example reference

For example, if your webhook request body looks like this:

```json
{
  "event": {
    "event_name": "event1",
    "data": {
      "value1": "value1",
      "value2": "value2"
    }
  }
}
```

Then your `./request/{name}_{UUID}.json` file will look like this:

```json
{
  "event.event_name": "event1",
  "event.data.value1": "value1",
  "event.data.value2": "value2"
}
```

> You can view the data keys of your webhooks in `./request/{name}_{UUID}.json`.

> Instructions for locating event data: If the event is included in headers, use `sickle.headers.{header name}` to
> locate the data,
> Otherwise, directly use the data key in `./request/{name}_{UUID}.json` to locate the data.


The following is a detailed description of the configuration file:

- `name` (string): The name of the webhook.
    - `url` (string): The generated webhook URL after running.
    - `source` (object): Specifies the source platform configuration for the webhook.
        - `type` (string): Specifies the type of source platform.
        - `config` (object): Configure details of the source platform, including event and target configuration.
            - `events` (array): A list containing one or more events. Each event object includes the following
              properties:
                - `name` (string): The name of the event.
                - `key` (string): the key of the event (in `./request/{name}_{UUID}.json`, used to locate the data of
                  the event, if the event is included in the headers,
                  use `sickle.headers.{header name}` to locate the data).
                - `destinations` (array): A list of destination platforms to send the event to when it fires.
    - `destinations` (array): A list containing one or more target platforms. Each target platform object includes the
      following properties:
        - `name` (string): The name of the target platform.
        - `type` (string): The type of target platform.
        - `config` (object): configuration information for the target platform, including webhook URL and custom data.
            - `webhook_url` (string): The webhook URL for the target platform.
            - `method` (string): The request method of the webhook.
            - `data` (array): A list containing one or more custom data. Each custom data object includes the following
              properties:
                - `event` (string): the name of the event (to facilitate the configuration of multiple events sharing a
                  target platform)
                - `data` (object): custom data object.
                    - `custom data name` (string): The value of the custom data. You can use `${incoming data key}` to
                      refer to the data of the source platform,
                      The data key of the specific source platform can be viewed in `./request/{name}_{UUID}.json`.
                    - `custom data name 2` (string): The value 2 of the custom data.


2. Configure on the web page

<mark>not yet implemented</mark>


## contribute

### Contribution template

```markdown
# Contributors

name: Pama-Lee

# Contributions

- [x] function 1
- [ ] function 2
```

### Contribution process

1. fork this repository
2. Clone the forked repository locally
3. Create a new branch
4. Make a change
5. Commit your changes
6. Push your changes to the remote repository
7. Create a pull request
8. Waiting for review

## License

[MIT](./LICENSE)