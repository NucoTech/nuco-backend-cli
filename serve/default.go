package serve

import (
	"fmt"
	"github.com/NucoTech/nuco-backend-cli/utils"
	"github.com/urfave/cli/v2"
	"net/http"
)

// 设置静态服务器
func setupStaticServe(dir string) error {
	fmt.Println("-----------------------------------------")
	fmt.Printf("\t静态服务已启动, 端口: %v\n", utils.StaticPort)
	fmt.Printf("\thttp://127.0.0.1:%v\n", utils.StaticPort)
	fmt.Println("-----------------------------------------")

	// 启动服务
	fmt.Println("服务监听中...")
	fmt.Println("使用 Ctrl/Command + C 即可退出...")
	http.Handle("/", http.FileServer(http.Dir(dir)))
	err := http.ListenAndServe(":" + utils.StaticPort, nil)

	return err
}

// 启动静态服务
func RegisterServeCommandAction() func(context *cli.Context) error {
	return func(context *cli.Context) error {
		dir := context.Args().Get(0)
		if len(dir) == 0 {
			dir = "docs"
		}
		if err := setupStaticServe(dir); err != nil {
			panic(err)
		}
		return nil
	}
}