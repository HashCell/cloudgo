package logging

// golang已经提供了log包，这里只是根据自定义的不同打印级别封装logger
import (
	"os"
	"log"
	"runtime"
	"fmt"
	"path/filepath"
)

type Level int

// iota出现在const语句块中的第几行，那么它就是几
// const中每新增一行常量声明将使iota计数一次
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// 全局变量
var (
	FHandle *os.File
	DefaultPrefix = ""
	DefaultCallerDepth = 2
	logger *log.Logger
	logPrefix = ""
	// levelFlags 的成员顺序一定要跟const枚举的顺序一致
	levelFlags = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
)

// 包init的时候，创建log文件
func init() {
	filePath := getLogFileFullPath()
	FHandle = openLogFile(filePath)
	// log.LstdFlags 定义log属性，打印log日期和时间
	// DefaultPrefix 定义每行log开头的前缀，创建的时候默认log前缀为空，调用时动态设置
	logger = log.New(FHandle, DefaultPrefix, log.LstdFlags)
}

// 调用不同级别的log打印，都需要先设置当前log打印的前缀，这是动态设置的
func Debug(v ...interface{}) {
	SetPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	SetPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	SetPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	SetPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	SetPrefix(FATAL)
	logger.Fatalln(v)
}
// 设置log打印的前缀
func SetPrefix(level Level) {
	// runtime.Caller returns program counter, file name, and line number within the file
	// here we only use filename, line
	_, AbsFilePath, lineNum, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		// define ours log prefix [Level] [filename : line number]
		// use filepath package to extract filename from absolute file path
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(AbsFilePath), lineNum)
	} else {
		// if not ok, we can only show the log level as prefix
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	// set prefix to logger
	logger.SetPrefix(logPrefix)
}