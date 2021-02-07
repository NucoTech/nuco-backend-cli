package docs

import (
	"bytes"
	"fmt"
	"github.com/NucoTech/nuco-backend-cli/utils"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

// 匹配title和href
var urlPattern = regexp.MustCompile(`<a class="js-navigation-open link-gray-dark" title="(.*?)" data-pjax="#repo-content-pjax-container" href="(.*?)">`)
var repositoryPattern = regexp.MustCompile(`(/.*?/.*?/)blob/(.*$)`)

const GITHUB = "https://github.com"

// 获取所有的文件链接下载
func download(client http.Client, url, dir string, wg *sync.WaitGroup) {
	if !utils.IsExist(dir){
		_ = os.Mkdir(dir, os.ModePerm)
	}
	// 获取文件页
	html, err := getHTML(client, url)
	if err != nil {
		fmt.Printf("获取模板页面失败: %s\n", err.Error())
		return
	}

	// 获取文件和目录
	links := urlPattern.FindAllSubmatch(html, -1)
	for _, link := range links {
		if isDir(link[2]) {
			download(client, GITHUB + string(link[2]), filepath.Join(dir, strings.SplitN(string(link[2]), "/", 6)[5]), wg)
		} else {
			rep := repositoryPattern.FindSubmatch(link[2])
			wg.Add(1)
			go utils.DownloadFile(client, utils.DocsTemplateBase + string(rep[2]), dir, string(link[1]), wg)
		}
	}
}

// 获取HTML
func getHTML(client http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	data , err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// 从模板生成文档
func generateDocsFromTemplate(dir string) error {
	var err error
	// 不在当前且不存在目录, 则默认创建目录
	if dir != "." && dir != "./" && !utils.IsExist(dir) {
		fmt.Printf("创建 %s 目录...\n", dir)
		err = os.Mkdir(dir, os.ModePerm)
	}

	var client http.Client
	var wg sync.WaitGroup
	// 开始时间
	start := time.Now()
	download(client, utils.DocsTemplateRepo, dir, &wg)
	wg.Wait()
	fmt.Printf("\n花费总时间: %.2f s!\n", float64(time.Since(start))/float64(time.Second))
	fmt.Printf("\n生成文档位于 %s 目录!!\n\nnbc serve\n\n即可启动文档预览!\n", dir)
	return err
}

// 判断是否是目录
func isDir(link []byte) bool {
	return bytes.Contains(link, []byte("tree"))
}

// 生成文档
func RegisterDocsCommandAction() func(context *cli.Context) error {
	return func(context *cli.Context) error {
		dir := context.Args().Get(0)
		if len(dir) == 0 {
			dir = "docs"
		}
		if err := generateDocsFromTemplate(dir); err != nil {
			panic(err)
		}
		return nil
	}
}
