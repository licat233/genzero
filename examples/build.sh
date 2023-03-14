#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd ${current_path}/../ || exit

if [ -f examples/genzero ]; then
    rm -f examples/genzero
fi

go mod tidy
go mod download
go build -o examples/genzero main.go
