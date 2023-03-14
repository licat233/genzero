#!/bin/bash
###
# @Author: licat
# @Date: 2023-01-11 15:40:07
 # @LastEditors: licat
 # @LastEditTime: 2023-02-17 22:59:19
# @Description: licat233@gmail.com
###

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

# ../build.sh

if [ ! -f "../genzero" ]; then
    ../build.sh
fi

# rm -f ./*.proto


if ! ../genzero pb --src="../sql/admin.sql" --service_name="admin" --multiple=false; then
    exit 1
fi

exit # 如果需要生成gozero框架的服务代码，请注释这行

if [ -d rpc ]; then
    rm -rf rpc/*
fi

# Please install the goctl first
# command: go install github.com/zeromicro/go-zero/tools/goctl@latest
# github: https://github.com/zeromicro/go-zero
goctl rpc protoc "admin.proto" --go_out="./rpc" --go-grpc_out="./rpc" --zrpc_out="./rpc"

cd ../
go mod tidy
