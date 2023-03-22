package logger

import (
	"fmt"
	"os"
)

func Debug(msg ...interface{}) {
	if os.Getenv("IS_DEV") == "true" {
		fmt.Println(msg...)
	}
}

func Info(msg ...interface{}) {
	fmt.Println(msg...)
}

func Warn(msg ...interface{}) {
	fmt.Println(msg...)
}

func Error(msg ...interface{}) {
	fmt.Println(msg...)
}
