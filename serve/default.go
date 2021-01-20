package serve

import (
	"fmt"
	"github.com/NucoTech/nuco-backend-cli/utils"
	"github.com/urfave/cli/v2"
	"net/http"
	"path/filepath"
)

func printInfo(port string, ssl bool) {
	onSSL := ""

	if ssl {
		onSSL = "s"
	}

	fmt.Println("-----------------------------------------")
	fmt.Printf("\t静态服务已启动, 端口: %v\n", port)
	fmt.Printf("\thttp%s://127.0.0.1:%v\n", onSSL, port)
	fmt.Println("-----------------------------------------")

	// 启动服务
	fmt.Println("服务监听中...")
	fmt.Println("使用 Ctrl/Command + C 即可退出...")
}

// 设置静态服务器
func setupStaticServe(dir, port string, ssl bool) error {
	var err error
	if !ssl {
		printInfo(port, ssl)
		http.Handle("/", http.FileServer(http.Dir(dir)))
		err = http.ListenAndServe(":" + port, nil)
	} else {
		// cert文件
		// key文件
		cert := filepath.Join(dir, "/cert.pem")
		key := filepath.Join(dir, "/key.pem")
		if utils.IsExist(cert) && utils.IsExist(key) {
			printInfo(port, ssl)
			http.Handle("/", http.FileServer(http.Dir(dir)))
			err = http.ListenAndServeTLS(":" + port, cert, key, nil)
		} else {
			fmt.Println("https静态服务未能启动成功!!!")
			fmt.Printf("密钥文件: %v 或证书文件: %v 不存在\n", key, cert)
			fmt.Println("请参阅nbc serve关于启动https静态服务部分的文档!!")
			err = nil
		}
	}

	return err
}

// 启动静态服务
func RegisterServeCommandAction() func(context *cli.Context) error {
	return func(context *cli.Context) error {
		dir := context.Args().Get(0)
		// 指定port
		port := context.String("port")

		ssl := context.Bool("S")

		if len(dir) == 0 {
			dir = "docs"
		}

		if len(port) == 0 {
			port = utils.DefaultStaticPort
		}

		if err := setupStaticServe(dir, port, ssl); err != nil {
			panic(err)
		}

		return nil
	}
}