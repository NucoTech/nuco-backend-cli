# 详解`nbc serve`

> `nbc serve`是一个本地启动静态服务的命令行工具

## 默认启动静态服务

```bash
nbc serve
```

> 启动执行目录的`/docs`子目录静态预览

## 执行目录启动静态服务

```bash
nbc serve .
```

## 任意目录启动静态服务

- 举个栗子, `666`子目录

```bash
nbc serve 666
```

## 指定端口

> nbc默认使用5001为静态服务的端口。在需要启动多静态服务的场景下, 可以通过指定端口号的方式启用多个服务

举个栗子, 使用`5002`端口启动`test`目录的静态服务

```bash
nbc serve -p 5002 test

# or

nbc serve --port 5002 test
```

## 支持TLS

> 自`1.3.x~`之后支持, nbc可以支持启动https服务, 因为在某些场景里有这个需求

- 生成本地ssl证书

> 生成证书需要依赖 `openssl` (Windows下请使用`Git Bash`终端, 由于系统目录结构的原因PowerShell会发生问题)

```bash
openssl genrsa -out key.pem 2048
openssl req -new -x509 -sha256 -key key.pem -out cert.pem
```

- 在执行目录启动服务

> 启动https服务简单, 只需要加上`-S`这个Flag

```bash
nbc serve -S .
```

## 默认目录`docs`预览支持TLS

> 由于`nbc`设计机制的缘故, 默认启动服务在`docs`文件夹下, 因此需要首先进入`docs`文件夹, 再遵守上述的步骤

- 进入文件夹

```bash
cd docs
```

- [遵循上述步骤](#支持TLS)
