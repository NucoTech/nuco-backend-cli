# 详解`nbc docs`

> `nbc docs`命令基于模板仓库[nuco-docsify](https://github.com/NucoTech/nuco-docsify)实现

## 默认生成文档

```bash
nbc docs
```

> 将会在执行目录下生成一个`/docs`子目录

## 执行目录生成文档

```bash
nbc docs .
```

## 任意目录生成文档

举个栗子, `666`目录

```bash
nbc docs 666
```

> 将会在执行目录下生成一个`/666`子目录

## 静态服务预览

> 请参考 [`nbc serve`](content/serve.md)详解
