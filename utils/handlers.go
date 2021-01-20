package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

func runCommand(cmd *exec.Cmd) {
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func RunGitAddCommand() {
	cmd := exec.Command("git", "add", ".")
	runCommand(cmd)
}

func RunGitCommitCommand(commit string)  {
	cmd := exec.Command("git", "commit", "-m", commit)
	runCommand(cmd)
}

func WriteFile(path, content string) {
	file, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0777)
	if err != nil {
		panic("文件打开错误")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic("文件关闭错误")
		}
	}()

	// 写入文件
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	if err != nil {
		panic("写入缓存错误")
	}
	_ = write.Flush()
}
func GetLineInput(toVar *string) (int, error) {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	*toVar = strings.TrimSpace(input)
	return len(*toVar), err
}

// 判断文件或者目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 下载文件
func DownloadFile(client http.Client, fileUrl, dir, filename string, wg *sync.WaitGroup) {
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

// 阻塞下载文件
func BlockDownloadFile(client http.Client, fileUrl, dir, filename string) {
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