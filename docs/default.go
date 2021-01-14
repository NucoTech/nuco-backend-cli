package docs

import (
	"bytes"
	"fmt"
	"github.com/NucoTech/nuco-backend-cli/utils"
	"github.com/urfave/cli/v2"
	"io"
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
var urlPattern = regexp.MustCompile(`<a class="js-navigation-open link-gray-dark" title="(.*?)" href="(.*?)">`)
var repositoryPattern = regexp.MustCompile(`(/.*?/.*?/)blob/(.*$)`)

const GITHUB = "https://github.com"

// 获取所有的文件链接下载
func download(client http.Client, url, dir string, wg *sync.WaitGroup) {
	if !isExist(dir){
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
			go downloadFile(client, utils.DocsTemplateBase + string(rep[2]), dir, string(link[1]), wg)
		}
	}
}

// 下载文件
func downloadFile(client http.Client, fileUrl, dir, filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("开始下载 %s...\n", filename)

	resp, err := client.Get(fileUrl)
	if err != nil {
		fmt.Printf("下载文件 %s 失败, 原因是: %s\n", filename, err.Error())
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var buff [1024]byte
	//	创建文件
	file, err := os.Create(filepath.Join(dir, filename))
	defer func() {
		_ = file.Close()
	}()

	if err != nil {
		fmt.Printf("创建文件 %s 错误\n", filename)
		return
	}

	// 文件写入
	for {
		n, err := resp.Body.Read(buff[:])
		_ , _ = file.Write(buff[:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("错误:", err)
			// 如果错误, 删掉这个文件
			_ = os.Remove(filepath.Join(dir, filename))
			return
		}
	}
	fmt.Printf("下载 %s 完成!\n", filename)
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
	if dir != "." && dir != "./" && !isExist(dir) {
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

// 判断文件或者目录是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 获取仓库名
func getRepositoryName(url string) string {
	var pattern *regexp.Regexp
	var count = strings.Count(url, "/")

	if count > 4{
		pattern = regexp.MustCompile(`https://github.com/.*?/(.*?)/`)
	}else if count == 4{
		pattern = regexp.MustCompile(`https://github.com/.*?/(.*$)`)
	}else{
		fmt.Println("url is wrong")
		os.Exit(-1)
	}
	name := pattern.FindStringSubmatch(url)
	return name[1]
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