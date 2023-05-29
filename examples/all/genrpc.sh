#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd "$(dirname $0)" || exit
    pwd
)

cd "$current_path" || exit

goctl rpc protoc admin.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --multiple=true --style=goZero
