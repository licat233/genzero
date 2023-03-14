/*
 * @Author: licat
 * @Date: 2023-02-18 10:35:19
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 11:31:36
 * @Description: licat233@gmail.com
 */

package tools

import (
	"fmt"
	"go/format"
)

// FormatGoContent 根据Go语言的语法规则自动格式化代码字符串
func FormatGoContent(code string) (string, error) {
	// 使用go/format包中的gofmt函数来格式化代码
	formatted, err := format.Source([]byte(code))
	if err != nil {
		return "", fmt.Errorf("go code format.Source: %w", err)
	}
	// 将格式化后的代码字符串返回
	return string(formatted), nil
}
