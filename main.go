package main

import (
	"os/exec"
	"strings"

	"github.com/licat233/genzero/cmd"
	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
)

func main() {
	out, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	if err != nil {
		tools.Error("获取git tags出错:%s", err)
		return
	}
	version := strings.TrimSpace(string(out))
	config.ProjectVersion = version
	cmd.Execute()
}
