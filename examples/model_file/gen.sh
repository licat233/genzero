#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0) || exit
    pwd
)

cd "$current_path" || exit

if [ ! -f "../genzero" ]; then
    ../build.sh
fi

# ../build.sh

# rm -rf ./model

if [ ! -d model ]; then
    if ! goctl model mysql ddl --src "../sql/admin.sql" --dir="model" --style goZero --cache true; then
        exit 1
    fi
    go mod tidy
fi

if ! ../genzero model --src="../sql/admin.sql" --service_name="admin" --dir="model";
then
    exit 1
fi