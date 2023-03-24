#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd "$(dirname $0)" || exit
    pwd
)

cd "$current_path" || exit

if [ -d model ]; then
    rm -rf model/*
fi

if [ -d api ]; then
    rm -rf api/*
fi

if [ -d rpc ]; then
    rm -rf rpc/*
fi

if ! goctl model mysql ddl --src "../sql/admin.sql" -dir="model" --style goZero -cache=false; then
    exit 1
fi

# Please install the goctl first
# command: go install github.com/zeromicro/go-zero/tools/goctl@latest
# github: https://github.com/zeromicro/go-zero

if ! goctl api go --api admin.api --dir ./api --style goZero; then
    exit 1
fi

if ! goctl rpc protoc admin.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --multiple=true --style goZero; then
    exit 1
fi

go mod tidy