#!/usr/bin/env bash

# 用于linux环境下交叉编译
echo "nbc项目交叉编译开始..."
echo ""
# Windows
export GOARCH="amd64" && GOOS="windows" && go build -o nbc.exe && echo "Windows平台编译完成!"
# MacOS
export GOARCH="amd64" && GOOS="darwin" && go build -o nbc.darwin && echo "MacOS平台编译完成!"
# MacOS arm64
export GOARCH="arm64" && GOOS="darwin" && go build -o nbc-arm64.darwin && echo "MacOS arm64平台编译完成!"
# Linux
export GOARCH="amd64" && GOOS="linux" && go build -o nbc.linux && echo "Linux平台编译完成!"
# Linux arm64
export GOARCH="arm64" && GOOS="linux" && go build -o nbc-arm64.linux && echo "Linux arm64平台编译完成!"
echo ""
echo "nbc项目交叉编译完成!!!"
