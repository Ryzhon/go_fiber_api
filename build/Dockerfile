FROM golang:1.20.4-alpine
# アップデートとgitのインストール！！
RUN apk add --update &&  apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
# # ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app
# RUN go env -w GO111MODULE=off

