# 使用するGolangのイメージを指定する
FROM golang:1.14-alpine

# Go Modules を有効化
ENV GO111MODULE=on

COPY go.mod .
# COPY go.sum .

# go.mod 内のパッケージをダウンロード
RUN go mod download

# ワーキングディレクトリを指定する
WORKDIR /go/src/api

#docker-compose.ymlの設定につながる
EXPOSE 8080
