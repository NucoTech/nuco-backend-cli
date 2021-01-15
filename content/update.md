# 更新`nbc`

> `nbc`目前不支持`nbc update`命令更新`nbc`(主要是懒得写emmmm), 请遵循下面的流程手动更新

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
