/*
 * @Author: licat
 * @Date: 2023-02-16 15:47:04
 * @LastEditors: licat
 * @LastEditTime: 2023-02-17 14:35:20
 * @Description: licat233@gmail.com
 */
package update

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
)

func Update() {
	currentVersion := config.CurrentVersion
	latestVersion := getLatestVersion()

	if !compareVersions(currentVersion, latestVersion) {
		tools.Success("当前版本%s已是最新.", currentVersion)
		return
	}

	tools.Success("新版本%s可用!", latestVersion)
	if err := updateSelf(latestVersion); err != nil {
		tools.Error("更新失败：%s\n", err)
		os.Exit(1)
	}
	tools.Success("更新成功，重新启动程序...")
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
	tools.Warning("正在更新: go install", url)
	// 构建并安装最新版本的程序
	if err := exec.Command(goBinary, "install", url).Run(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func compareVersions(oldVer, newVer string) bool {
	oldParts := strings.Split(oldVer, ".")
	newParts := strings.Split(newVer, ".")

	for i := 0; i < len(newParts) && i < len(oldParts); i++ {
		oldNum, oldStr := parseVersionPart(oldParts[i])
		newNum, newStr := parseVersionPart(newParts[i])

		if newNum > oldNum {
			return true
		} else if newNum < oldNum {
			return false
		}

		// If both parts are equal, but one has a suffix and the other doesn't,
		// consider the one without the suffix to be newer.
		if oldStr == "" && newStr != "" {
			return true
		} else if oldStr != "" && newStr == "" {
			return false
		}
	}

	// If all parts match, but the second version has more parts, it's considered newer.
	return len(newParts) > len(oldParts)
}

// Parses a single part of a version number, which may contain numbers and/or letters.
// Returns the numerical part as an integer, and the non-numerical part as a string.
func parseVersionPart(part string) (int, string) {
	dotIdx := strings.IndexByte(part, '.')
	if dotIdx >= 0 {
		num, _ := strconv.Atoi(part[0:dotIdx])
		return num, part[dotIdx+1:]
	}

	num, err := strconv.Atoi(part)
	if err == nil {
		// If the whole part is a number, return it as-is.
		return num, ""
	}

	// Otherwise, split the part into numeric and non-numeric parts.
	for i := len(part); i > 0; i-- {
		s := part[:i]
		if digit, err := strconv.Atoi(s); err == nil {
			return digit, ""
		}
	}

	// If there are no numeric characters in the part, treat it as 0.
	return 0, part
}
