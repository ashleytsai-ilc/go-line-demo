# Go Line Demo

> 使用Golang, Gin及MongoDB的入門練習

實作串接Linebot，以官方帳號收發訊息，並將訊息存在MongoDB

專案使用項目:

* [Gin](https://github.com/gin-gonic/gin)
* [Viper](https://github.com/spf13/viper)
* [Cobra](https://github.com/spf13/cobra)
* [Mongo Driver](https://github.com/mongodb/mongo-go-driver)
* [Linebot SDK Go](https://github.com/line/line-bot-sdk-go)

使用docker image:
* [golang:alpin](https://hub.docker.com/_/golang/) as base image
* [mongo:4.4](https://hub.docker.com/_/mongo)

### 專案目錄結構

```
.
├── Config
│   └── config.go
├── controllers
│   └── message.go
├── database
│   └── mongo.go
├── models
│   └── message.go
├── routes
│   ├── index.go
│   └── message.go
├── utils
│   └── linebot.go
├── validators
│   ├── message.go
│   └── register.go
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── README.md

```

### 執行專案
#### 初始化
* Step1. 建立環境變數檔案
```sh
cp .env.example .env
```
* Step2. 修改.env參數內容，包含MongoDB連線參數及Linebot相關密鑰設定
```sh
DB_CONNECTION = mongodb
DB_HOST = mongo
DB_PORT = 27017
DB_USERNAME = admin
DB_PASS = admin

LINE_CHANNEL_ID =
LINE_CHANNEL_SECRET =
LINE_CHANNEL_ACCESS =
```
* Step3. 登入[Line Developers](https://developers.line.biz/en/)
* Step4. 點擊[Messaging API](https://developers.line.biz/en/services/messaging-api/)進行官方帳號及Webhook設定
將URL設定到Webhook

  *PS. local端測試使用[Ngrok](https://dashboard.ngrok.com/get-started/setup)*

  `https://{some domain}/messages/receive`

#### 啟動專案
```sh
docker-compose up -d --build
```

### 使用範例
* 連接測試

`curl https://{some domain}`
* 取得官方帳號訊息列表

`[GET] https://{some domain}/user/{Line official account user id}/messages`
* 利用官方帳號推送訊息

`[POST] https://{some domain}/messages/push`
```sh
# Content-Type: application/json
# Body parameters
{
    "UserID": "Line official account user id",
    "Content": "some message you want to push"
}
```
