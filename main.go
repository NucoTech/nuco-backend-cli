package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"tech.nuco.nbc/commit"
)

const (
	Version = "1.0.0"
)

// TODO 注册命令行工具命令
func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "commit",
				Usage:  "生成标准commit提交信息",
				Action: commit.RegisterCommitCommandAction(),
			},
			{
				Name: "version",
				Usage: "打印版本信息",
				Action: func(context *cli.Context) error {
					fmt.Println("工具仓库地址: https://github.com/NucoTech/nuco-backend-cli")
					fmt.Printf("当前工具的版本号为:\t%s\n", Version)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
