# Cinnox Line Bot 
[![Build Status][github-action-status]][github-action-url]

Linebot service API implemented the following features
- Get messages that a user send in a period of time.
- Send message to a certain user with line user ID.

## Table of Contents
---
  - [Let's Try](#lets-try)
  - [Packages](#packages)
  - [Commit Style](#commit-style)
  - [Run in Local](#run-in-local)
  - [API](#api)
    - [Get message](#get-message)
      - [Request params](#request-params)
      - [Response body](#response-body)
    - [Push message to a certain user](#push-message-to-a-certain-user)
      - [Request body](#request-body)
      - [Response body](#response-body-1)

## Let's Try

- Join the line bot
  - Scan the QRCode to join

  ![](./line-bot-qrcode.png)

  - Use Bot ID to join: @574haccu

- Send message to yourself with API

- Get the message you have sent to the bot
  - This [page](https://www.epochconverter.com/) would help you transform timestamp to unix time

```shell
curl 'https://cinnox.nekosekai.com/message?user={your-user-id}&start_time=1&end_time=1676002045'
```

- Send some message to the bot

```shell
curl -X POST 'https://cinnox.nekosekai.com/message' \
-H 'Content-Type: application/json' \
-d '{
    "user":"{your-user-id}",
    "content":"Hello"
}'
```

> **Warning**
> You could only use the API above after sending a message to this bot

> **Note**
> Just in case you don't want to add this bot to your line account, 
> you could query my data

```shell
curl 'https://cinnox.nekosekai.com/message?user=U731727d29c9f4944fee1f0a4987acf35&start_time=1&end_time=1676002045'
```

```shell
curl -X POST 'https://cinnox.nekosekai.com/message' \
-H 'Content-Type: application/json' \
-d '{
    "user":"U731727d29c9f4944fee1f0a4987acf35",
    "content":"Hello from Cinnox!"
}'
```

## Packages
- [gin](https://github.com/gin-gonic/gin) for rest service
- [zap](https://github.com/uber-go/zap) for logging
- [viper](https://github.com/spf13/viper) for setting up config
- [cobra](https://github.com/spf13/cobra) for building command line
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) for mongoDB operations
- [line-bot-sdk](https://github.com/line/line-bot-sdk-go) for linebot client



## Commit Style
---
- Follow [Conventional Commit Messages](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13)

## Run in Local
---
- Add user config in {project_dir}/config/config.yaml
- Build the app with `go build`


```shell
go build .
```

- Start mongoDB service with `Makefile`
  - note that the Makefile would auto install docker cli and colima runtime
    
```shell
make dev
```

- Run the command generated by Cobra to initialize the collection `message.line`

```shell
./{your-app-name} create line
```

- Start the service

```shell
./{your-app-name}
```

- Expose Service with [Ngrok](https://ngrok.com/)
  - You could follow this [article](https://ngrok.com/docs/getting-started) for ngrok installation and usage

```shell
ngrok http {your-service-port}
```

- Parse Ngrok Host to LineBot Webhook URL
   - You could follow this [document](https://developers.line.biz/en/docs/messaging-api/building-bot/) for the linebot quick setup

## API
---

### Get message 

#### Request params
| Field | Type | Description |  
| :---- | :----| :---------- | 
| user  | string| user to query | 
| start_time | int64 | unix time | 
| end_time | int64 | unix time | 

#### Response body
| Field | Type | Description |  
| :---- | :----| :---------- | 
| status | string | Success or Failed | 
| user  | string| user to query  | 
| time | int64 | unix micro time | 
| content | string | message text content | 

- request sample
```shell
curl 'http://{your-host}:{your-port}/message?user=U731727d29c9f4944fee1f0a4987acf35&start_time=1&end_time=1676004680'
```

--- 


### Push message to a certain user

#### Request body
| Field | Type | Description |  
| :---- | :----| :---------- | 
| user  | string| user to send the message| 
| content | string | message content | 

#### Response body
| Field | Type | Description |  
| :---- | :----| :---------- | 
| status | string | Success or Failed | 
| user  | string| user to send | 
| content | string | message content | 

- request sample
```shell
curl -X POST 'http://{your-host}:{your-port}/message' \
-H 'Content-Type: application/json' \
-d '{
    "user":"U731727d29c9f4944fee1f0a4987acf35",
    "content":"Hello"
}'
```

[github-action-status]: https://github.com/yusianglin11010/cinnox-line-bot/workflows/deploy/badge.svg?branch=main
[github-action-url]: https://github.com/yusianglin11010/cinnox-line-bot/actions?query=branch
