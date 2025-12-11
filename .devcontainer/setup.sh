#!/bin/bash
set -e
# bash-completionのインストール(単に補完を使いたかっただけ)
sudo apt update && sudo apt install bash-completion
echo '. /usr/share/bash-completion/completions/git' >> ~/.bashrc

# ロケールをUTF-8に設定
echo 'export LANG=C.UTF-8' >> ~/.bashrc
echo 'export LC_ALL=C.UTF-8' >> ~/.bashrc

# bashrcを読み込み
source ~/.bashrc

# Git設定
echo "Setting up Git configuration..."
git config --global --add safe.directory /app

# 環境変数のチェック
if [ -z "$GIT_USER_EMAIL" ] || [ -z "$GIT_USER_NAME" ]; then
    echo "Error: GIT_USER_EMAIL and GIT_USER_NAME must be set"
    exit 1
fi

git config --global user.email "$GIT_USER_EMAIL"
git config --global user.name "$GIT_USER_NAME"

# Go依存関係のダウンロード
echo "Downloading Go dependencies..."
cd /app/src
go mod download

echo "Setup completed!"