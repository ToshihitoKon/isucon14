#!/usr/bin/env bash
cd "$(dirname "$0")"
git fetch
git merge origin/main

# 設定ファイルをコピー
sudo cp -r /nginx /etc/
sudo cp -r /mysql /etc/

# アプリケーションの再起動
sudo nginx -t
sudo systemctl restart mysql
sudo systemctl restart nginx

echo "デプロイ"
