package tools

import "fmt"

func Message(format string, args ...interface{}) (n int, err error) {
	return fmt.Printf(format+"\n", args...)
}

func Warning(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf("\033[33m%s\033[0m\n", fmt.Sprintf(format, a...))
}

func Success(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf("\033[32m%s\033[0m\n", fmt.Sprintf(format, a...))
}

func Error(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf("\033[31m%s\033[0m\n", fmt.Sprintf(format, a...))
}
