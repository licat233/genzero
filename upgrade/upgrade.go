package upgrade

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
)

func Upgrade() {
	currentVersion := config.ProjectVersion
	latestVersion := getLatestVersion()

	newest := isNewestVersion(currentVersion, latestVersion)

	if strings.EqualFold(currentVersion, latestVersion) || newest {
		tools.Success("当前版本[%s]已是最新.", currentVersion)
		return
	}

	tools.Success("新版本[%s]可用!", latestVersion)
	if err := updateSelf(latestVersion); err != nil {
		tools.Error("更新失败：%s\n", err)
		os.Exit(1)
	}
	tools.Success("更新成功!")
	// 重新启动程序
	os.Exit(0)
}

// 获取最新版本号
func getLatestVersion() string {
	v, err := tools.GetLatestReleaseVersion(config.ProjectInfoURL)
	if err != nil {
		tools.Error("获取最新版本号失败：%s\n", err)
		os.Exit(1)
	}
	return v
}

// 自我升级
func updateSelf(latestVersion string) error {
	goBinary, err := exec.LookPath("go")
	if err != nil {
		return fmt.Errorf("未找到go命令：%w", err)
	}
	url := strings.ReplaceAll(config.ProjectURL, "http://", "")
	url = strings.ReplaceAll(url, "https://", "") + "@" + latestVersion
	tools.Success("更新命令: go install %s\n正在更新...", url)
	// 构建并安装最新版本的程序
	if err := exec.Command(goBinary, "install", url).Run(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func isNewestVersion(version1, version2 string) bool {
	return compareVersions(version1, version2) >= 0
}

// 比较两个版本，如果version1大于version2，则返回1，否则返回-1，相等返回0
func compareVersions(version1, version2 string) int {
	// 版本号相同，直接返回 0
	if strings.EqualFold(version1, version2) {
		return 0
	}

	// 去除前缀 "v"，并删除预发布版本标识
	version1 = strings.TrimPrefix(version1, "v")
	arr1 := strings.Split(version1, "-")
	version1a := arr1[0]
	version2 = strings.TrimPrefix(version2, "v")
	arr2 := strings.Split(version2, "-")
	version2a := arr2[0]

	if !strings.EqualFold(version1a, version2a) {
		// 分割字符串为版本号数组
		v1 := strings.Split(version1a, ".")
		v2 := strings.Split(version2a, ".")

		// 比较每个部分的大小
		for i := 0; i < len(v1) || i < len(v2); i++ {
			var num1, num2 int

			// 转换版本号部分为数字
			if i < len(v1) {
				num1, _ = strconv.Atoi(v1[i])
			}

			if i < len(v2) {
				num2, _ = strconv.Atoi(v2[i])
			}

			// 比较当前部分的大小
			if num1 > num2 {
				return 1
			} else if num1 < num2 {
				return -1
			}
		}
	}

	//如果比较下来，"-"符号前都相同，则比较"-"，后半部分
	var str1, str2 string
	if len(arr1) > 1 {
		str1 = arr1[1]
	}
	if len(arr2) > 1 {
		str2 = arr2[1]
	}

	if str1 == "" {
		if str2 != "" {
			return 1
		} else {
			return 0
		}
	} else {
		if str2 != "" {
			arr1 = strings.Split(str1, ".")
			arr2 = strings.Split(str2, ".")
			s1 := arr1[0]
			s2 := arr2[0]
			if s1 == s2 {
				//第一部分相同，比较最后一部分
				s1a := arr1[len(arr1)-1]
				s2a := arr2[len(arr2)-1]
				num1, _ := strconv.Atoi(s1a)
				num2, _ := strconv.Atoi(s2a)
				// 比较当前部分的大小
				if num1 > num2 {
					return 1
				} else if num1 < num2 {
					return -1
				}
			} else if s1 == "bate" {
				if s2 == "alpha" {
					return 1
				} else {
					return -1
				}
			} else if s2 == "bate" {
				if s1 == "alpha" {
					return -1
				} else {
					return 1
				}
			}
		} else {
			return -1
		}
	}

	// 版本号相等
	return 0
}
