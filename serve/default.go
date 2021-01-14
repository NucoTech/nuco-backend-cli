package serve

import (
	"fmt"
	"github.com/NucoTech/nuco-backend-cli/utils"
	"github.com/urfave/cli/v2"
	"net/http"
)

// 设置静态服务器
func setupStaticServe(dir, port string) error {
	fmt.Println("-----------------------------------------")
	fmt.Printf("\t静态服务已启动, 端口: %v\n", port)
	fmt.Printf("\thttp://127.0.0.1:%v\n", port)
	fmt.Println("-----------------------------------------")

	// 启动服务
	fmt.Println("服务监听中...")
	fmt.Println("使用 Ctrl/Command + C 即可退出...")
	http.Handle("/", http.FileServer(http.Dir(dir)))
	err := http.ListenAndServe(":" + port, nil)

	return err
}

// 启动静态服务
func RegisterServeCommandAction() func(context *cli.Context) error {
	return func(context *cli.Context) error {
		dir := context.Args().Get(0)
		// 指定port
		port := context.String("port")

		if len(dir) == 0 {
			dir = "docs"
		}

		if len(port) == 0 {
			port = utils.DefaultStaticPort
		}

		if err := setupStaticServe(dir, port); err != nil {
			panic(err)
		}

		return nil
	}
}