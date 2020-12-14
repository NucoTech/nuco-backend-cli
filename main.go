package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"tech.nuco.nbc/commit"
	"tech.nuco.nbc/initProj"
	"tech.nuco.nbc/utils"
)

// TODO 注册命令行工具命令
func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "new",
				Usage: "初始化工程配置文件",
				Action: initProj.RegisterInitCommandAction(),
			},
			{
				Name:   "commit",
				Usage:  "生成标准commit提交信息",
				Action: commit.RegisterCommitCommandAction(),
			},
			{
				Name: "info",
				Usage: "打印工具信息",
				Action: func(context *cli.Context) error {
					fmt.Println("工具仓库地址: https://github.com/NucoTech/nuco-backend-cli")
					fmt.Printf("当前版本:\t%s\n", utils.VERSION)
					return nil
				},
			},
			{
				Name: "version",
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
