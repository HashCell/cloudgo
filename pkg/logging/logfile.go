package logging

import (
	"fmt"
	"time"
	"os"
	"log"
)

// log文件扩展名　.log
var (
	LogSavePath = "runtime/logs"
	LogSaveName = "log"
	LOgFileExt = "log"
	TImeFormat = "20200418"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TImeFormat),LOgFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// switch的用法跟其他语言有点区别
func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	// 判断文件是否存在，不存在则创建
	case os.IsNotExist(err):
		mkDir(getLogFilePath())
		// 如果系统调用权限出错
	case os.IsPermission(err):
		log.Fatalf("Permission : %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to openfile :%v", err)
	}
	return handle
}

func mkDir(filePath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir + "/" + filePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
