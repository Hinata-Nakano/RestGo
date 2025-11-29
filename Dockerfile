# goバージョン
FROM golang:tip-trixie
# ワーキングディレクトリの設定
WORKDIR /app
#ディレクトリのコピー
COPY ./src .
# パッケージの管理
RUN go mod download && go mod tidy