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
	// 去除前缀 "v"
	version1 = strings.TrimPrefix(version1, "v")
	version2 = strings.TrimPrefix(version2, "v")

	// 分割字符串为版本号数组
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

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

	// 版本号相等
	return 0
}
