package main

import (
	"fmt"
	"github.com/NucoTech/nuco-backend-cli/commit"
	"github.com/NucoTech/nuco-backend-cli/docs"
	"github.com/NucoTech/nuco-backend-cli/initProj"
	"github.com/NucoTech/nuco-backend-cli/serve"
	"github.com/NucoTech/nuco-backend-cli/update"
	"github.com/NucoTech/nuco-backend-cli/utils"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "初始化工程配置文件",
				Action: initProj.RegisterInitCommandAction(),
			},
			{
				Name:   "commit",
				Usage:  "生成标准commit提交信息",
				Action: commit.RegisterCommitCommandAction(),
			},
			{
				Name:   "docs",
				Usage:  "生成文档模板",
				Action: docs.RegisterDocsCommandAction(),
			},
			{
				Name:   "serve",
				Usage:  "启动静态服务",
				Action: serve.RegisterServeCommandAction(),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "指定端口号 (default: 5001)",
					},
					&cli.BoolFlag{
						Aliases: []string{"S"},
						Usage:   "启动https服务",
					},
				},
			},
			{
				Name:   "update",
				Usage:  "更新nbc工具",
				Action: update.RegisterUpdateCommandAction(),
			},
			{
				Name:  "info",
				Usage: "打印工具信息",
				Action: func(context *cli.Context) error {
					fmt.Println("工具仓库地址: https://github.com/NucoTech/nuco-backend-cli")
					fmt.Printf("当前版本:\t%s\n", utils.VERSION)
					return nil
				},
			},
			{
				Name:  "version",
				Usage: "打印版本信息",
				Action: func(context *cli.Context) error {
					fmt.Println(utils.VERSION)
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
