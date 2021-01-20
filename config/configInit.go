package config

import (
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

var configDefault = []byte(`
[commit]
text = [
    "feat :sparkles: ",
    "fix :bug: ",
    "docs :pencil: ",
    "perf :zap: ",
    "test :white_check_mark: ",
    "chore :wrench: ",
    "refactor :recycle: ",
    "revert :rewind: ",
    "release :bookmark: ",
    "deploy :rocket: ",
    "ci :construction_worker: ",
]
description = [
    "✨ feat: 新的特性",
    "🐛 fix: 修复bug",
    "📝 docs: 更改文档",
    "⚡️ perf: 提升性能",
    "✅ test: 代码测试",
    "🔧 chore: 项目配置相关",
    "♻️ refactor: 代码重构",
    "⏪ revert: 回滚提交",
    "🔖 release: 发布新版本",
    "🚀 deploy: 项目部署",
    "👷 ci: 持续集成",
]
`)

func exportConfigFile() {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(homeDir, ".nbcrc")
	file, err := os.Create(path)
	if err != nil {
		panic(errors.New("创建配置文件失败"))
	}
	_, err = file.Write(configDefault)
	if err != nil {
		panic(errors.New("配置文件写入失败"))
	}
	defer func() {
		_ = file.Close()
		fmt.Printf("\n配置文件导出成功!\n路径为:\t%v\n", path)
	}()
}

func RegisterConfigInitCommandAction() func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		fmt.Println("nbc config init调用")
		exportConfigFile()
		return nil
	}
}
