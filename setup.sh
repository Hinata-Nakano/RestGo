#!/bin/sh

# hookの参照先を.git/hooksから.githooksに変更
git config --local core.hooksPath .githooks

# hookを実行可能にする
chmod +x .githooks/*

echo "✓ Git hooksパスが設定されました: .githooks"
echo "  設定を確認: git config --local --list | grep core.hookspath"