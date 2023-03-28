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

rm -f ./*.api
rm -f ./*.proto

# ../build.sh

# if [ ! -f "./genzeroConfig.yaml" ]; then
#     if ! ../genzero init config; then
#         exit 1
#     fi
# fi

# 会根据yaml配置文件来生成服务的配置文件
if ! ../genzero start --src="./config.yaml" --dev; then
    exit 1
fi