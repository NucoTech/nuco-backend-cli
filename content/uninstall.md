# 卸载`nbc`

## Windows平台

1. 直接删除掉`nbc.exe`的所在目录
2. 删除环境变量`PATH`中的上述字段

## 类Unix平台(含Linux、MacOS)

- 脚本卸载

```bash
wget https://cdn.jsdelivr.net/gh/NucoTech/nuco-backend-cli@main/uninstall.sh
bash uninstall.sh
```

- 手动卸载

```bash
sudo rm -f /usr/bin/nbc
```
