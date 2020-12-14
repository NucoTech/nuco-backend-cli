#!/usr/bin/env bash

print_nbc_version() {
    nbc version
}

if [ -e /usr/bin/nbc ]
then
    # 本地存在nbc
    echo ""
    echo "本地已存在nbc命令行工具!"
    echo ">>> 路径位于 /usr/bin/nbc"
    echo ""
    print_nbc_version
else
    echo ">>> 正在使用curl下载nbc..."
    wget https://github.com/NucoTech/nuco-backend-cli/releases/latest/download/nbc.linux && sudo chmod +x nbc.linux && sudo mv nbc.linux /usr/bin/nbc
    print_nbc_version
fi
