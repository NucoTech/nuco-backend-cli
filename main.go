package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"tech.nuco.nbc/commit"
)

// TODO 注册命令行工具命令
func main() {
	//t := prompt.Input(">>> ", commit.CommitTypeCompleter)
	//fmt.Print(t)
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "commit",
				Usage:  "生成标准commit提交信息",
				Action: commit.RegisterCommitCommandAction(),
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
