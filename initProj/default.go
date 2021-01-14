package initProj

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/NucoTech/nuco-backend-cli/utils"
)

// 生成.gitignore文件
func generateGitIgnoreFile() {
	fmt.Println("生成.gitignore文件中...")
	filePath := ".gitignore"
	fileContent := ".idea/\n"
	utils.WriteFile(filePath, fileContent)
	fmt.Println(".gitignore生成成功！")
}

// 注册初始化命令行
func RegisterInitCommandAction() func(context *cli.Context) error {
	return func(context *cli.Context) error {
		generateGitIgnoreFile()
		// 执行git commit
		utils.RunGitAddCommand()
		utils.RunGitCommitCommand(":tada: Initial Goland Project")
		return nil
	}
}