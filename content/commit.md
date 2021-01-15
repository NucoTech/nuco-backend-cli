# 详解`nbc commit`

> `nbc commit`一个是为了取代`git commit -m`的命令

## 标准提交流程

- 使用`nbc commit`之前

```bash
# 添加git管理
git add .

# 添加描述
git commit -m ""

# push代码
git push
```

- 使用`nbc commit`之后

```bash
# 添加git管理
git add .

# 进入nbc commit交互
nbc commit

# push代码
git push
```

## `nbc commit`交互流程

> 以实际交互为准

```text
? 选择你提交的类型
? 本次提交的简述
? 本次提交的具体描述
? 本次提交是否存在 BREAKING CHANGES (不兼容更新)
? 本次提交是否关闭已知的issue (eg. Fix #1 #2)
? 是否确定本次提交
```

## 标准commit规范

```text
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

- type

> `type`字段基于angular规范进行业务情况拓展, emoji风格

| 类型     | 说明             |
| -------- | :--------------- |
| feat     | 新的特性         |
| fix      | bug修复          |
| docs     | 文档相关         |
| pref     | 性能提升         |
| test     | 代码测试         |
| chore    | 项目配置相关     |
| refactor | 代码重构         |
| revert   | 撤销上一次commit |
| release  | 发布新版本       |
| deploy   | 项目部署         |
| ci       | 持续集成工具     |

- scope

> 修改的范围, `nbc commit`不支持该字段

- subject

> 提交的简短描述

- body

> 提交信息的正文

- footer

> 补充信息，是否含 `BREAKING CHANGES` 和 关闭issues?

## `nbc commit`自定义规范

`nbc commit`采用了绘文字(emoji)化的commit风格, 为了让commit更加有辨识度并且**可愛い**

> 所有使用的emoji都遵循规范, 参见 [gitmoji](https://gitmoji.dev/)

`nbc commit`支持的`commit`类型:

```go
[]prompt.Suggest{
    {Text: "feat :sparkles: ", Description: "✨ feat: 新的特性"},
    {Text: "fix :bug: ", Description: "🐛 fix: 修复bug"},
    {Text: "docs :pencil: ", Description: "📝 docs: 更改文档"},
    {Text: "perf :zap: ", Description: "⚡️ perf: 提升性能"},
    {Text: "test :white_check_mark: ", Description: "✅ test: 代码测试"},
    {Text: "chore :wrench: ", Description: "🔧 chore: 项目配置相关"},
    {Text: "refactor :recycle: ", Description: "♻️ refactor: 代码重构"},
    {Text: "revert :rewind: ", Description: "⏪ revert: 回滚提交"},
    {Text: "release :bookmark: ", Description: "🔖 release: 发布新版本"},
    {Text: "deploy :rocket: ", Description: "🚀 deploy: 项目部署"},
    {Text: "ci :construction_worker: ", Description: "👷 ci: 持续集成"},
}
```

## What's More

> 目前`nbc commit`没有支持用户自定义commit规范的计划, 可能会在之后的某个版本支持通过配置文件加载commit自定义规范
