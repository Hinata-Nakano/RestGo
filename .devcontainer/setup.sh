#!/bin/bash
set -e

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