package config

import (
	"github.com/NucoTech/nuco-backend-cli/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

// 检查配置文件是否存在
func CheckIfConfigFileExist() bool {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return utils.IsExist(filepath.Join(homeDir, ".nbcrc"))
}

// 读取配置文件
func ReadConfigFile() {
	//homeDir, err := homedir.Dir()
	//if err != nil {
	//	panic(err)
	//}
	// 检查配置文件是否存在
}

func RegisterConfigCommandAction() func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		ReadConfigFile()
		return nil
	}
}
