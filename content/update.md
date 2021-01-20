# 更新`nbc`

## 命令行更新

> 自`v1.3.x~`之后支持`nbc update`更新。由于权限问题, 请遵循提示的步骤完成更新!

## Windows平台

1. 进入`nbc.exe`所在的目录
2. 在[Release](https://github.com/NucoTech/nuco-backend-cli/releases)中找到最新发布的`nbc.exe`
3. 下载替换

## Linux平台

- 如果你之前下载了安装脚本和卸载脚本

```bash
bash uninstall.sh
bash install.linux.sh
```

- 手动更新

```bash
sudo rm -f /usr/bin/nbc

wget https://github.com/NucoTech/nuco-backend-cli/releases/latest/download/nbc.linux
sudo chmod +x nbc.linux
sudo mv nbc.linux /usr/bin/nbc
```

## MacOS平台

- 如果你之前下载了安装脚本和卸载脚本

```bash
bash uninstall.sh
bash install.darwin.sh
```

- 手动更新

```bash
sudo rm -f /usr/bin/nbc

wget https://github.com/NucoTech/nuco-backend-cli/releases/latest/download/nbc.darwin
sudo chmod +x nbc.darwin
sudo mv nbc.darwin /usr/bin/nbc
```
