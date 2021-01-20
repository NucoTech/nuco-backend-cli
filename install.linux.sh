#!/usr/bin/env bash

print_nbc_info() {
    nbc info
}

if [ -e /usr/bin/nbc ]
then
    # 本地存在nbc
    echo ""
    echo "本地已存在nbc命令行工具!"
    echo ">>> 路径位于 /usr/bin/nbc"
    echo ""
    print_nbc_info
else
    echo ">>> 正在使用wget下载nbc..."
    wget https://github.xiu2.xyz/https://github.com/NucoTech/nuco-backend-cli/releases/latest/download/nbc.linux && sudo chmod +x nbc.linux && sudo mv nbc.linux /usr/bin/nbc
    print_nbc_info
fi
