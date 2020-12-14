package utils

import (
	"bufio"
	"os"
	"os/exec"
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

func WriteFile(path, content string)  {
	file, err := os.OpenFile(path, os.O_WRONLY, 0777)
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
	_, err =write.WriteString(content)
	if err != nil {
		panic("写入缓存错误")
	}
	err = write.Flush()
	if err != nil {
		panic("写入文件错误")
	}
}