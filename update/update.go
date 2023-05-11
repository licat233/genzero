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
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
)

func Update() {
	currentVersion := config.CurrentVersion
	latestVersion := getLatestVersion()

	if strings.EqualFold(currentVersion, latestVersion) {
		tools.Success("当前版本%s已是最新.", currentVersion)
		return
	}

	tools.Success("新版本%s可用!", latestVersion)
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
	tools.Success("正在更新...")
	// 构建并安装最新版本的程序
	if err := exec.Command(goBinary, "install", url).Run(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// func compareVersions(oldVer, newVer string) bool {
// 	oldVer, newVer = strings.TrimSpace(oldVer), strings.TrimSpace(newVer)
// 	oldLen, newLen := len(oldVer), len(newVer)
// 	minLen := oldLen
// 	if oldLen > newLen {
// 		minLen = newLen
// 	}

// 	var oldVerScore, newVerScore int
// 	for i := 0; i < minLen; i++ {
// 		oldV, newV := oldVer[i], newVer[i]
// 		if oldV == newV {
// 			continue
// 		}
// 		oldDigit, oldErr := strconv.Atoi(string(oldV))
// 		newDigit, newErr := strconv.Atoi(string(newV))
// 		if oldErr != nil && newErr == nil {
// 			return true
// 		}
// 		if newErr != nil && oldErr == nil {
// 			return false
// 		}
// 		if oldErr != nil && newErr != nil {
// 			return true
// 		}
// 		if oldErr == nil && newErr == nil {
// 			if newDigit > oldDigit {
// 				newVerScore++
// 			} else if newDigit < oldDigit {
// 				oldVerScore++
// 			} else {
// 				continue
// 			}
// 		}
// 	}
// 	if oldVerScore == newVerScore {
// 		return oldLen > newLen
// 	}

// 	return oldVerScore < newVerScore
// }
