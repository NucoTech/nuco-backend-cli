package commit

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/urfave/cli/v2"
	"tech.nuco.nbc/utils"
)

var (
	validCommit = [11]string{
"feat :sparkles: ",
"fix :bug: ",
"docs :pencil: ",
"perf :zap: ",
"test :white_check_mark: ",
"chore :wrench: ",
"refactor :recycle: ",
"revert :rewind: ",
"release :bookmark: ",
"deploy :rocket: ",
"ci :construction_worker: ",
}
)

func commitTypeCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "feat :sparkles: ", Description: "✨ feat: 新的特性"},
		{Text: "fix :bug: ", Description: "🐛 fix: 修复bug"},
		{Text: "docs :pencil: ", Description: "📝 docs: 更改文档"},
		{Text: "perf :zap: ", Description: "⚡️ perf: 提升性能"},
		{Text: "test :white_check_mark: ", Description: "✅ test: 代码测试"},
		{Text: "chore :wrench: ", Description: "🔧 chore: 项目配置相关"},
		{Text: "refactor :recycle: ", Description: "♻️ refactor: 代码重构"},
		{Text: "revert :rewind: ", Description: "⏪ revert: 回滚提交"},
		{Text: "release :bookmark: ", Description: "🔖 release: 发布新版本"},
		{Text: "deploy :rocket: ", Description: "🚀 deploy: 项目部署"},
		{Text: "ci :construction_worker: ", Description: "👷 ci: 持续集成"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func getCommitInput(structType string, handler *string) {
	fmt.Print(">>> ")
	n, err := utils.GetLineInput(handler)

	if err != nil {
		panic("获取输入错误")
	}

	switch structType {
	case "subject": {
		if n == 0 {
			panic("简述的值不得为空")
		}
	}
	case "body":
	case "issue": {
		if n == 0 {
			*handler = ""
		}
	}
	case "BREAKING": {
		if *handler == "Y" {
			*handler = "BREAKING CHANGES"
		} else {
			*handler = ""
		}
	}
	case "verified":
		if !(n == 0 || *handler == "Y") {
			panic("已经取消本次提交")
		} else {
			*handler = "Y"
		}
	}
}

func checkCommitType(rawValue string) bool {
	for i, v := range validCommit {
		if i == len(validCommit) && v != rawValue {
			return false
		}

		if v != rawValue {
			continue
		} else {
			return true
		}
	}
	return false
}

func makeCommit(commitType, commitSubject, commitBody, commitBroken, commitIssues string) string {
	return fmt.Sprintf("%s: %s\n\n%s\n\n%s\n", commitType, commitSubject, commitBody, commitBroken + commitIssues)
}

func RegisterCommitCommandAction() func(ctx *cli.Context) error {
	return func(context *cli.Context) error {
		var commitSubject string
		var commitBody string
		var commitBroken string
		var commitIssues string
		var commitVerified string
		fmt.Println("? 选择你提交的类型 (Tab键自动填充)")
		commitType := prompt.Input(">>> ", commitTypeCompleter)
		// 检查提交
		if !checkCommitType(commitType) {
			panic("提交的type不合格")
		}

		// 获取输入
		fmt.Println("? 本次提交的简述")
		getCommitInput("subject", &commitSubject)
		fmt.Println("? 本次提交的具体描述 (可选)")
		getCommitInput("body", &commitBody)
		fmt.Println("? 本次提交是否存在 BREAKING CHANGES (不兼容更新, Y/n 默认为 n)")
		getCommitInput("BREAKING", &commitBroken)
		fmt.Println("? 本次提交是否关闭已知的issue (可选, eg. #1 #2)")
		getCommitInput("issue",  &commitIssues)

		// 生成commit
		commit := makeCommit(commitType, commitSubject, commitBody, commitBroken, commitIssues)
		fmt.Println(commitType, commitSubject, commitBody, commitBroken, commitIssues)
		fmt.Printf("\n本次生成的提交信息:\n%s\n", commit)
		fmt.Println("? 是否确定本次提交 (Y/n, 默认为Y)")
		getCommitInput("verified",  &commitVerified)
		if commitVerified == "Y" {
			utils.RunGitCommitCommand(commit)
			fmt.Println("commit已生成, 可以使用git push提交")
		}
		return nil
	}
}