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
	fmt.Printf("当前版本号：%s\n", currentVersion)
	latestVersion := getLatestVersion()
	fmt.Printf("最新版本号：%s\n", latestVersion)
	if latestVersion != currentVersion {
		fmt.Println("有新版本可用，开始更新...")
		if err := updateSelf("latest"); err != nil {
			fmt.Printf("更新失败：%s\n", err)
			os.Exit(1)
		}
		fmt.Println("更新成功，重新启动程序...")
		// 重新启动程序
		os.Exit(0)
	} else {
		fmt.Println("已是最新版本，无需更新。")
	}
}

// 获取最新版本号
func getLatestVersion() string {
	v, err := tools.GetLatestReleaseVersion(config.ProjectInfoURL)
	if err != nil {
		fmt.Printf("获取最新版本号失败：%s\n", err)
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
	// 构建并安装最新版本的程序
	if err := exec.Command(goBinary, "install", url).Run(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
