# nuco-backend-cli

nuco后端命令行工具，含commit规范化提交等

## 安装方式

1. 在release中找到最新的发行版
2. 根据不同的平台下载对应的命令行工具
3. 将命令行工具的放置位置加入系统环境变量 (or PATH)

## 注意事项

- Windows平台下请使用cmd而不是powershell，powershell的emoji有显示问题

## 可用命令

| 命令 | 说明 |
| :--- | :--- |
| `nbc commit` | 启动规范化commit交互 |

## commit交互流程

```shell
? 选择你提交的类型

? 本次提交的简述

? 本次提交的具体描述

? 本次提交是否存在 BREAKING CHANGES (不兼容更新)

? 本次提交是否关闭已知的issue (eg. #1 #2)

? 是否确定本次提交
```

## commit指令

### 使用`nbc commit`命令替代`git commit -m`

```shell
git add .
nbc commit
git push
```

### 标准commit规范

```text
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

- type

> `type`字段基于angular规范进行业务情况拓展, emoji风格

| 类型 | 说明 |
| --- | :--- |
| feat | 新的特性 |
| fix | bug修复 |
| docs | 文档相关 |
| pref | 性能提升 |
| test | 代码测试 |
| chore | 项目配置相关 |
| refactor | 代码重构 |
| revert | 撤销上一次commit |
| release | 发布新版本 |
| deploy | 项目部署 |
| ci | 持续集成工具 |

- scope

> 修改的范围, 默认不启用这一层

- subject

> 提交的简短描述

- body

> 提交信息的正文

- footer

> 补充信息，是否含 `BREAKING CHANGES` 和 关闭issues?
