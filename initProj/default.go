package initProj

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"tech.nuco.nbc/utils"
)

// 生成.gitignore文件
func generateGitIgnoreFile() {
	filePath := ".gitignore"
	fileContent := ".idea/\n"
	utils.WriteFile(filePath, fileContent)
}

// 注册初始化命令行
func RegisterInitCommandAction() func(context *cli.Context) error {
	fmt.Println("生成.gitignore文件...")
	generateGitIgnoreFile()
	// 执行git commit
	utils.RunGitAddCommand()
	utils.RunGitCommitCommand(":tada: Initial Goland Project")
	return nil
}