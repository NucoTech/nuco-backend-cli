# nuco-backend-cli

nuco后端命令行工具，含commit规范化提交等

## 安装方式

1. 在release中找到最新的发行版
2. 根据不同的平台下载对应的命令行工具
3. 将命令行工具的放置位置加入系统环境变量 (or PATH)

> `nbc.exe`为Windows版, `nbc.darwin`和`nbc.linux`分别为MacOS和Linux版, 下载之后务必重命名

```shell
# MacOS
mv nbc.darwin nbc

# Linux
mv nbc.linux nbc
```

## 注意事项

- Windows平台下请使用cmd而不是powershell，powershell的emoji有显示问题

## 可用命令

| 命令 | 说明 |
| :--- | :--- |
| `nbc commit` | 启动规范化commit交互 |

## 文档

- [commit命令指南](./docs/commit命令使用指南.md)

## TODOs

- [ ] 支持`upgrade`命令升级
- [ ] 支持`init`命令生成通用配置文件
- [ ] 需要重写`version`指令
