#!/usr/bin/env bash
set -eux
cd "$(dirname "$0")"
rootdir=`pwd`

# git fetch
# git merge origin/main

cd "${rootdir}/go"
go build -o isuride ./...

cd $rootdir
# 設定ファイルをコピー
sudo cp -r ./etc/nginx /etc/
sudo cp -r ./etc/mysql /etc/

# アプリケーションの再起動
sudo nginx -t
sudo systemctl restart mysql
sudo systemctl restart nginx

echo "デプロイ"
~
~
