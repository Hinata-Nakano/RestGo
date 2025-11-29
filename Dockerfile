# goバージョン
FROM golang:tip-trixie
#ディレクトリのコピー
COPY ./src ./app
# パッケージの管理
RUN go mod download && go mod tidy