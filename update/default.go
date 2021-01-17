package update

import (
	"fmt"
	"github.com/NucoTech/nuco-backend-cli/utils"
	"github.com/urfave/cli/v2"
	"regexp"
	"strconv"
)

var versionPattern = regexp.MustCompile(`v([0-9]*).([0-9]*).([0-9]*)`)

const repo = "https://github.com/NucoTech/nuco-backend-cli"

// 获取远端最新版本
func getRemoteVersion() (string, error) {
	//res, err := http.Get(repo)
	//if err != nil {
	//	return "", err
	//}

	return "", nil
}

// 比较版本号
func compareVersion()  {
	localVersion, err := getVersion(utils.VERSION)
	fmt.Println(localVersion)
	if err != nil {
		panic(err)
	}
}

// 版本号解析
func getVersion(version string) ([]int, error) {
	versionBytes := versionPattern.FindSubmatch([]byte(version))
	version1, err := strconv.Atoi(string(versionBytes[1]))
	version2, err := strconv.Atoi(string(versionBytes[2]))
	version3, err := strconv.Atoi(string(versionBytes[3]))
	if err != nil {
		return nil, err
	}
	return []int{version1, version2, version3}, nil
}

func RegisterUpdateCommandAction() func(context *cli.Context) error {
	return func(context *cli.Context) error {
		compareVersion()
		return nil
	}
}
