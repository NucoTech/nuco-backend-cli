# nuco-backend-cli

nuco后端命令行工具，含commit规范化提交等

## 安装方式

### 类Unix平台

- `MacOS`

```shell
wget https://github.com/NucoTech/nuco-backend-cli/releases/latest/download/install.darwin.sh
bash install.darwin.sh
```

- `Linux`

```shell
wget https://github.com/NucoTech/nuco-backend-cli/releases/latest/download/install.linux.sh
bash install.linux.sh
```

### Windows平台

- 在`release`中找到最新发布的`nbc.exe`
- 将下载的`nbc.exe`所在目录加入环境变量PATH

## 注意事项

- Windows平台下请使用cmd而不是powershell，powershell的emoji有显示问题

## 可用命令

| 命令 | 说明 |
| :--- | :--- |
| `nbc commit` | 启动规范化commit交互 |

## 文档

- [commit命令指南](commit命令使用指南.md)
