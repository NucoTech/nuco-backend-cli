#!/usr/bin/env bash

# 用于linux环境下交叉编译
echo "nbc项目交叉编译开始..."
echo ""
# Windows
export GOOS="windows" && go build -o nbc.exe && echo "Windows平台编译完成!"
# MacOS
export GOOS="darwin" && go build -o nbc.darwin && echo "MacOS平台编译完成!"
# Linux
export GOOS="linux" && go build -o nbc.linux && echo "Linux平台编译完成!"
echo ""
echo "nbc项目交叉编译完成!!!"