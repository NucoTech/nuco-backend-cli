# 安装 Install

`nbc`支持Windows、Linux、MacOS这些不同平台的使用, 跟随下面不同平台的教程可以让你快速安装`nbc`工具

## Windows平台

1. 在[Release](https://github.com/NucoTech/nuco-backend-cli/releases)中找到最新发布的`nbc.exe`
2. 将下载的`nbc.exe`存放的地址加入环境变量`PATH`

## Linux平台

- 使用脚本安装

```bash
wget https://cdn.jsdelivr.net/gh/NucoTech/nuco-backend-cli@main/install.linux.sh
bash install.linux.sh
```

- 手动安装

```bash
wget https://github.com/NucoTech/nuco-backend-cli/releases/latest/download/nbc.linux
sudo chmod +x nbc.linux
sudo mv nbc.linux /usr/bin/nbc
```

## MacOS平台

- 使用脚本安装

```bash
wget https://cdn.jsdelivr.net/gh/NucoTech/nuco-backend-cli@main/install.darwin.sh
bash install.darwin.sh
```

- 手动安装

```bash
wget https://github.com/NucoTech/nuco-backend-cli/releases/latest/download/nbc.darwin
sudo chmod +x nbc.darwin
sudo mv nbc.darwin /usr/bin/nbc
```
