#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 '{version}'"
    exit 1
fi

# build binary
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o miniprogram -ldflags "-X main.version=$1"

if [ $? -eq 0 ]; then
    echo "build success"
    docker build -t zhoushoujian/dingtalk_webhook:latest .
    if [ $? -eq 0 ]; then
        echo "docker build success"
        docker push zhoushoujian/dingtalk_webhook:latest
    else
        echo "docker build failed"
    fi
else
    echo "build failed"
fi



