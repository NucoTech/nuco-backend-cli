package update

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NucoTech/nuco-backend-cli/utils"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type LatestReleaseType struct {
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	AssetsUrl string `json:"assets_url"`
	UploadUrl string `json:"upload_url"`
	TarballUrl string `json:"tarball_url"`
	ZipballUrl string `json:"zipball_url"`
	Id int `json:"id"`
	NodeId string `json:"node_id"`
	TagName string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
	Name string `json:"name"`
	Body string `json:"body"`
	Draft bool `json:"draft"`
	Prerelease bool `json:"prerelease"`
	CreatedAt string `json:"created_at"`
	PublishedAt string `json:"published_at"`
	Author interface{} `json:"author"`
	Assets []interface{} `json:"assets"`
}

var versionPattern = regexp.MustCompile(`v([0-9]*).([0-9]*).([0-9]*)`)

const repo = "https://api.github.com/repos/NucoTech/nuco-backend-cli/releases/latest"

// 新版本检查
func ifLatestVersion(now, remote []int) bool {
	for i := 0; i < len(now); i++ {
		if remote[i] > now[i] {
			return false
		} else if remote[i] < now[i] {
			return true
		}
	}
	return true
}

// 获取远端最新版本
func getRemoteVersion() (string, error) {
	res, err := http.Get(repo)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	// 数据结构体
	var latestRelease LatestReleaseType
	if err := json.Unmarshal(body, &latestRelease); err != nil {
		return "", err
	}

	return latestRelease.TagName, nil
}

// 比较版本号
func CompareLatestVersion() (string, bool) {
	localVersion, err := getVersion(utils.VERSION)
	remote, err := getRemoteVersion()
	remoteVersion, err := getVersion(remote)

	// 错误处理
	if err != nil {
		panic(err)
	}

	if !ifLatestVersion(localVersion, remoteVersion) {
		return remote, false
	} else {
		return utils.VERSION, true
	}
}

// 获取对应平台软件
func backPlatformVersion(version string) (url, filename string) {
	url = utils.NewVersionCDN + version + "/nbc."
	filename = "nbc."
	switch runtime.GOOS {
		case "windows": {
			url = url + "exe"
			filename = filename + "exe.tmp"
		}
		case "darwin": {
			url = url + "darwin"
			filename = filename + "darwin.tmp"
		}
		case "linux": {
			url = url + "linux"
			filename = filename + "linux.tmp"
		}
	}
	return filename, url
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

func getHomeDir() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func RegisterUpdateCommandAction() func(context *cli.Context) error {
	return func(context *cli.Context) error {
		version, latest := CompareLatestVersion()
		filename, url := backPlatformVersion(version)

		if !latest {
			var c http.Client

			if runtime.GOOS == "windows" {
				// 下载到当前文件夹
				if utils.IsExist("nbc.exe") {
					dir, _ := os.Getwd()
					utils.BlockDownloadFile(c, url, dir, filename)
					// 删除旧版本
					fmt.Println("\n请手动删除nbc.exe, 并将nbc.exe.tmp文件重命名为nbc.exe以完成更新")
					fmt.Println("\n执行 nbc info 即可查看更新后的信息")
				} else {
					panic(errors.New("请在nbc.exe对应的文件夹执行更新"))
				}
			} else {
				home, err := getHomeDir()
				if err != nil {
					panic(err)
				}
				utils.BlockDownloadFile(c, url, home, filename)

				cmd := exec.Command("chmod", "+x", filepath.Join(home, filename))
				err = cmd.Run()
				if err != nil {
					panic(err)
				}
				fmt.Printf("授予 %v 可执行权限成功!\n", filepath.Join(home, filename))
				fmt.Printf("\n请手动执行 sudo mv %v /usr/bin/nbc 以更新完成!", filepath.Join(home, filename))
				fmt.Println("\n执行 nbc info 即可查看更新后的信息")
			}
		} else {
			fmt.Printf("当前版本为最新版本, 版本号为 %v\n", version)
		}

		return nil
	}
}
