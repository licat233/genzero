/*
 * @Author: licat
 * @Date: 2023-01-14 11:12:42
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 10:38:38
 * @Description: licat233@gmail.com
 */
package tools

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

// 名词复数形式，不完全准
func PluralizedName(name string) string {
	chip := name[len(name)-1:]
	switch chip {
	case "s":
		return name + "list"
		// if name[len(name)-2:] == "es" || name[len(name)-2:] == "ys" {
		// 	return name + "list"
		// }
		// return name + "es"
	case "y":
		if name[len(name)-2:] == "ey" {
			return name + "s"
		}
		return name[:len(name)-1] + "ies"
	case "_":
		return name[:len(name)-1] + "list"
	default:
		return name + "s"
	}
}

func StringEqual(a, b string) bool {
	return ToSnake(a) == ToSnake(b)
}

// 不区分大小写查找
func SliceContain(slice []string, s string) bool {
	s = ToSnake(s)
	for _, v := range slice {
		// if strings.EqualFold(v, s) {
		// 	return true
		// }
		if ToSnake(v) == s {
			return true
		}
	}

	return false
}

// func InSlice(target string, sources ...string) bool {
// 	return HasInSlice(sources, target)
// }

func PickMarkContents2(startMark, endMark string, content []byte) ([][]byte, error) {
	if len(content) == 0 {
		return [][]byte{}, nil
	}
	expr := fmt.Sprintf("%s((?s).*?)%s", startMark, endMark)

	reg, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	listArr := reg.FindAllSubmatch(content, -1)
	byteArr := [][]byte{}
	for _, list := range listArr {
		target := list[len(list)-1]
		if len(target) == 0 {
			continue
		}
		byteArr = append(byteArr, target)
	}
	// Message("%s匹配结果: %#v \n", expr, byteArr)
	return byteArr, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

var FileExists = PathExists

func MakeDir(filename string) error {
	err := os.MkdirAll(filename, 0755)
	return err
}

func RTCFile(filename string) (content string, f *os.File, err error) {
	//read
	fd, e := os.ReadFile(filename)
	if e != nil && !errors.Is(e, os.ErrNotExist) {
		err = e
		return
	}

	//to string
	content = string(fd)
	n := len(content)
	content = strings.TrimSpace(content)
	content = strings.Trim(content, "\n")
	content = strings.Trim(content, "\t")
	for n != len(content) {
		content = strings.TrimSpace(content)
		content = strings.Trim(content, "\n")
		content = strings.Trim(content, "\t")
		n = len(content)
	}

	//读写 | 清空 | 创建
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)

	return
}

func GetFilename(filename string) string {
	// 获取文件名
	filename = filepath.Base(filename)

	// 获取文件类型
	extension := filepath.Ext(filename)

	return strings.TrimSuffix(filename, extension)
}

func SetFileType(filepath, filetype string) string {
	fileType := path.Ext(filepath)
	if fileType != filetype {
		filename := filepath[0 : len(filepath)-len(fileType)]
		filepath = fmt.Sprintf("%s%s", filename, filetype)
	}
	return filepath
}

func FileRename(oldFilepath, newname string) string {
	// 获取文件所在目录
	directory := filepath.Dir(oldFilepath)

	// 获取文件名
	// filename := filepath.Base(oldFilepath)

	// 获取文件类型
	extension := filepath.Ext(oldFilepath)

	// filetype := extension[1:]

	newFilename := fmt.Sprintf("%s%s", newname, extension)
	return filepath.Join(directory, newFilename)
}

func ToCamel(s string) string {
	return strcase.ToCamel(s)
}

func ToLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

func ToSnake(s string) string {
	return strcase.ToSnake(strcase.ToCamel(s))
}

func IsSnake(s string) bool {
	// 定义正则表达式
	r := regexp.MustCompile(`^[a-z]+(_[a-z]+)*$`)
	return r.MatchString(s)
}

func IsCamel(s string) bool {
	// 定义正则表达式
	r := regexp.MustCompile(`^[A-Z]+([A-Z][a-z]*)*$`)
	return r.MatchString(s)
}

func IsLowerCamel(s string) bool {
	// 定义正则表达式
	r := regexp.MustCompile(`^[a-z]+([A-Z][a-z]*)*$`)
	return r.MatchString(s)
}

func ExecShell(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

// 调用shell升级go版本
func UpgradeCurrentProject(currentVersion, projectInfoURL, projectUrl string) error {
	//获取最新版本号
	version, err := GetLatestReleaseVersion(projectInfoURL)
	if err != nil {
		return fmt.Errorf("\n * Failed to get the latest version number，project info api url: %s \n   error: %s", projectInfoURL, version)
	}
	//对比版本号
	if version == currentVersion {
		Success(" version: %s\n The current version is the latest version，no need to upgrade，\n", currentVersion)
		return nil
	}

	//先检查go是否存在
	if _, err = exec.LookPath("go"); err != nil {
		//不存在，则提示先安装go
		return errors.New("\n * warning: \n   go not exist\n   Please install go first")
	}

	// 运行shell命令，调用go install进行升级
	url := strings.ReplaceAll(projectUrl, "http://", "")
	url = strings.ReplaceAll(url, "https://", "")
	command := fmt.Sprintf("go install %s@latest", url)
	if out, err := ExecShell(command); err != nil {
		return errors.New(out)
	}

	Success(" Upgrade succeeded: %s -> %s\n", currentVersion, version)
	return nil
}

// 获取github项目的最新release版本号
func GetLatestReleaseVersion(projectInfoURL string) (string, error) {
	command := fmt.Sprintf("wget -qO- -t1 -T2 \"%s\" | grep \"tag_name\" | head -n 1 | awk -F \":\" '{print $2}' | sed 's/\\\"//g;s/,//g;s/ //g'", projectInfoURL)
	out, err := ExecShell(command)
	out = strings.TrimSpace(out)
	return out, err
}

// 获取git用户名
func GetGitUserName() (string, error) {
	cmd := exec.Command("git", "config", "user.name")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	username := strings.TrimSpace(string(out))
	return username, nil
}

// 获取git用户邮箱
func GetGitUserEmail() (string, error) {
	cmd := exec.Command("git", "config", "user.email")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	username := strings.TrimSpace(string(out))
	return username, nil
}

// 获取系统用户名
func GetOsUserName() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.Username, err
}

// 获取当前用户名
func GetCurrentUserName() string {
	author, err := GetGitUserName()
	if err != nil || author == "" {
		author, _ = GetOsUserName()
	}
	return author
}

func GetCurrentDirectory() (string, error) {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1), nil
}

func GetCurrentDirectoryName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	names := strings.Split(dir, "/")
	return names[len(names)-1], nil
}

func FindFilename(dir string, file string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, fileInfo := range files {
		if strings.EqualFold(fileInfo.Name(), file) {
			return fileInfo.Name(), nil
		}
	}
	return "", nil
}

// 查找文件，以忽略文件名大小写的方式查找匹配
func FindFile(dir string, file string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, fileInfo := range files {
		if strings.EqualFold(fileInfo.Name(), file) {
			return filepath.Join(dir, fileInfo.Name()), nil
		}
	}
	return "", nil
}

func CreateFile(filePath string) error {
	// 获取文件所在的目录
	dir := filepath.Dir(filePath)

	// 创建目录及其父目录
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 创建文件
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	// 关闭文件
	defer func(f *os.File) {
		if f != nil {
			err := f.Close()
			if err != nil {
				return
			}
		}
	}(f)

	return nil
}

func WriteFile(filename string, data string) error {
	exists, err := FileExists(filename)
	if err != nil {
		return err
	}
	if !exists {
		err := CreateFile(filename)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(filename string) (string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return TrimSpace(string(b)), nil
}

func TrimSpace(content string) string {
	n := len(content)
	content = strings.TrimSpace(content)
	content = strings.Trim(content, "\n")
	content = strings.Trim(content, "\t")
	for n != len(content) {
		content = strings.TrimSpace(content)
		content = strings.Trim(content, "\n")
		content = strings.Trim(content, "\t")
		n = len(content)
	}
	return content
}

func IsIdColumn(name string) bool {
	// name = ToSnake(name)
	// names := strings.Split(name, "_")
	// n := len(names)
	// if n == 0 {
	// 	return false
	// }
	// return names[n-1] == "id"
	return strings.HasSuffix(name, "_id") || strings.HasSuffix(name, "Id")
}

func ExecGoimports(w string) error {
	// 检查命令是否存在
	goimportsBinary, err := exec.LookPath("goimports")
	if err != nil {
		if err := InstallGoImports(); err != nil {
			return err
		}
	}
	if w == "" {
		w = "."
	}
	// 命令存在，执行它
	err = exec.Command(goimportsBinary, "-w", w).Run()
	if err != nil {
		return fmt.Errorf("command failed to run: gofmt -w %s", w)
	}
	return nil
}

func InstallGoImports() error {
	goBinary, err := exec.LookPath("go")
	if err != nil {
		return fmt.Errorf("未找到go命令：%w", err)
	}
	err = exec.Command(goBinary, "install", "golang.org/x/tools/cmd/goimports@latest").Run()
	if err != nil {
		return errors.New("failed to install command: go install golang.org/x/tools/cmd/goimports@latest，Please install manually")
	}
	return nil
}

func ExecGoFormat(filename string) error {
	// 检查命令是否存在
	goimportsBinary, err := exec.LookPath("gofmt")
	if err != nil {
		return fmt.Errorf("未找到gofmt命令：%w\n请先安装go：https://go.dev", err)
	}
	if filename == "" {
		filename = "."
	}
	// 命令存在，执行它
	err = exec.Command(goimportsBinary, "-w", filename).Run()
	if err != nil {
		return fmt.Errorf("command failed to run: gofmt -w %s", filename)
	}
	return nil
}

func Template(name string) *template.Template {
	return template.New(name).Funcs(template.FuncMap{
		"ToCamel":      ToCamel,
		"ToLowerCamel": ToLowerCamel,
		"ToSnake":      ToSnake,
	})
}

func IsReturn(line string) bool {
	line = strings.TrimSpace(line)
	reg := regexp.MustCompile(`^return\s.+{},\snil$`)
	return reg.MatchString(line)
}

func ParserTpl(tpl string, data any) (string, error) {
	var buf bytes.Buffer
	t, err := Template("tpl").Parse(tpl)
	if err != nil {
		return "", err
	}
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func FormatGoFile(filename string) (err error) {
	filename, err = GetGoFileDir(filename)
	if err != nil {
		return err
	}

	if err := ExecGoimports(filename); err != nil {
		return err
	}

	if err := ExecGoFormat(filename); err != nil {
		return err
	}
	return nil
}

func GetGoFileDir(filename string) (string, error) {
	if !strings.HasSuffix(filename, ".go") {
		return filename, nil
	}
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		return "", err
	}
	return dir, nil
}

func PickGoPkgName(line string) string {
	line = strings.TrimSpace(line)
	if line == "" {
		return ""
	}
	reg := regexp.MustCompile(`(\w+)$`)
	arr := reg.FindStringSubmatch(line)
	if len(arr) == 2 {
		return strings.TrimSpace(arr[1])
	}
	return ""
}

func IsTimeTypeField(field string) bool {
	field = strings.TrimSpace(field)
	parts := []string{
		"Time",
		"_time",
		"At",
		"_at",
	}
	for i := range parts {
		if strings.HasSuffix(field, parts[i]) {
			return true
		}
	}
	return false
}
