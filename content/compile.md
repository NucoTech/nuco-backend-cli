# 手动编译

> 自 `v1.3.2` 开始, `nbc` 支持Apple M1 和 Linux/arm64 (eg. [Termux](https://termux.com/), [AidLearning](http://www.aidlearning.net/)) 平台。但因编译arm64平台发行版比较麻烦, 在目前的 `nbc` 发行版中无编译完成的releases, 可以自行编译使用。

请注意 **Apple M1** 平台编译支持需要将 Golang 升级到 **1.16**

## Windows平台

```bash
export GOARCH="amd64" && GOOS="windows" && go build -o nbc.exe
```

## Linux平台

```bash
export GOARCH="amd64" && GOOS="linux" && go build -o nbc
```

- arm64

```bash
export GOARCH="arm64" && GOOS="linux" && go build -o nbc
```

## MacOS平台

```bash
export GOARCH="amd64" && GOOS="darwin" && go build -o nbc
```

- arm64

```bash
export GOARCH="arm64" && GOOS="darwin" && go build -o nbc
```
