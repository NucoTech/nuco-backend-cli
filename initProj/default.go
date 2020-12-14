package initProj

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"tech.nuco.nbc/utils"
)

// 生成.gitignore文件
func generateGitIgnoreFile() {
	fmt.Println("生成.gitignore文件中...")
	filePath := ".gitignore"
	fileContent := ".idea/\n"
	utils.WriteFile(filePath, fileContent)
}

// 注册初始化命令行
func RegisterInitCommandAction() func(context *cli.Context) error {
	generateGitIgnoreFile()
	// 执行git commit
	utils.RunGitAddCommand()
	utils.RunGitCommitCommand(":tada: Initial Goland Project")
	return nil
}