#!/bin/bash

# CGO_ENABLED 跨平台时禁用
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server-linux .  # 当前项目根目录
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./server .
CGO_ENABLED=0 GOOS=linux GOARCH=windows go build -o ./server.exe .