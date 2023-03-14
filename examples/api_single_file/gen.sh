#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

if [ ! -f "../genzero" ]; then
    ../build.sh
fi

# ../build.sh

# rm -f ./*.api

if ! ../genzero api --src="../sql/admin.sql" --service_name="admin-api" --jwt="Auth" --middleware="AuthMiddleware" --prefix="/v1/api/admin" --multiple=false --ignore_tables="jwt_blacklist"; then
    exit 1
fi

exit


if [ -d api ]; then
    rm -rf api/*
fi

# Please install the goctl first
# command: go install github.com/zeromicro/go-zero/tools/goctl@latest
# github: https://github.com/zeromicro/go-zero

# exit # 如果需要生成gozero框架的服务代码，请注释这行

# goctl template init --home ./template

if ! goctl api go -api admin.api -dir ./api -style goZero --home ./template; then
    exit 1
fi

cd ../
go mod tidy
