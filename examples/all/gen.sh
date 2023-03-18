#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd "$(dirname $0)" || exit
    pwd
)

cd "$current_path" || exit

if [ ! -f "../genzero" ]; then
    ../build.sh
fi

if [ -d model ]; then
    rm -rf model/*
fi

if [ -d api ]; then
    rm -rf api/*
fi

if [ -d rpc ]; then
    rm -rf rpc/*
fi

# ../build.sh

if ! goctl model mysql ddl --src "../sql/admin.sql" -dir="model" --style goZero -cache=false; then
    exit 1
fi


if [ ! -f "./genzeroConfig.yaml" ]; then
    if ! ../genzero init; then
        exit 1
    fi
fi

# 会根据yaml配置文件来生成服务的配置文件
if ! ../genzero yaml; then
    exit 1
fi

# exit # 如果需要生成gozero框架的服务代码，请注释这行

# Please install the goctl first
# command: go install github.com/zeromicro/go-zero/tools/goctl@latest
# github: https://github.com/zeromicro/go-zero

if ! goctl api go --api admin.api --dir ./api --style goZero; then
    exit 1
fi

if ! goctl rpc protoc admin.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --multiple=true; then
    exit 1
fi

go mod tidy
