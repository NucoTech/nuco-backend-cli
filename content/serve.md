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
