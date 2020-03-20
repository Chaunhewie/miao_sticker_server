package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/phachon/go-logger"
)

type MyLogger struct {
	Path   string
	logger *go_logger.Logger
}

var myLogger *MyLogger = &MyLogger{}
var myContactLogger *MyLogger = &MyLogger{}

func init() {
	myLogger.logger = go_logger.NewLogger()
	myLogger.Path = strings.Replace(os.Getenv("GOPATH"), "\\", "/", -1) + "/src/miao_sticker_server/index/output/app.log"
	myLogger.logger.SetAsync()
	myLogger.registerConsole()
	myLogger.registerFile()

	myContactLogger.logger = go_logger.NewLogger()
	myContactLogger.Path = strings.Replace(os.Getenv("GOPATH"), "\\", "/", -1) + "/src/miao_sticker_server/index/output/contact.log"
	myContactLogger.logger.SetAsync()
	myContactLogger.registerConsole()
	myContactLogger.registerContactFile()
}

func (myLogger *MyLogger) registerConsole() {
	logger := myLogger.logger

	if err := logger.Detach("console"); err != nil {
		fmt.Println("Logger Detach console Error:", err)
	}
	// 命令行输出配置
	consoleConfig := &go_logger.ConsoleConfig{
		Color:      true,  // 命令行输出字符串是否显示颜色
		JsonFormat: false, // 命令行输出字符串是否格式化
		Format:     "%level_string% %millisecond_format% %body%",
	}
	// 添加 console 为 Logger 的一个输出
	if err := logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig); err != nil {
		fmt.Println("Logger Attach console Error:", err)
	}
}

func (myLogger *MyLogger) registerFile() {
	logger := myLogger.logger

	// 文件输出配置
	fileConfig := &go_logger.FileConfig{
		Filename:   myLogger.Path, // 日志输出文件名，不自动存在
		MaxSize:    1024 * 1024,   // 文件最大值（KB），默认值0不限
		MaxLine:    0,             // 文件最大行数，默认 0 不限制
		DateSlice:  "d",           // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: false,         // 写入文件的数据是否 json 格式化
		Format:     "%level_string% %millisecond_format% %body%",
	}
	// 添加 file 为 Logger 的一个输出
	if err := logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig); err != nil {
		fmt.Println("Logger Attach file Error:", err)
	}
}

func (myLogger *MyLogger) registerContactFile() {
	logger := myLogger.logger

	// 文件输出配置
	fileConfig := &go_logger.FileConfig{
		Filename:   myLogger.Path, // 日志输出文件名，不自动存在
		MaxSize:    0,             // 文件最大值（KB），默认值0不限
		MaxLine:    100000,        // 文件最大行数，默认 0 不限制
		DateSlice:  "m",           // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: false,         // 写入文件的数据是否 json 格式化
		Format:     "%level_string% %millisecond_format% %body%",
	}
	// 添加 file 为 Logger 的一个输出
	if err := logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig); err != nil {
		fmt.Println("Logger Attach file Error:", err)
	}
}

func getLogBasicInfo() (fileName string, funcName string, line int) {
	funcName = "null"
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "null"
		line = 0
	} else {
		funcName = runtime.FuncForPC(pc).Name()
	}
	_, fileName = path.Split(file)
	return
}

func Debug(format string, a ...interface{}) {
	if myLogger.logger != nil {
		filename, funcName, line := getLogBasicInfo()
		dataStr := fmt.Sprintf(format, a...)

		myLogger.logger.Debug(fmt.Sprintf("%v %v:%v data=[%v]", filename, funcName, line, dataStr))
	}
}

func Info(format string, a ... interface{}) {
	if myLogger.logger != nil {
		filename, funcName, line := getLogBasicInfo()
		dataStr := fmt.Sprintf(format, a...)

		myLogger.logger.Info(fmt.Sprintf("%v %v:%v data=[%v]", filename, funcName, line, dataStr))
	}
}

func Warning(format string, a ... interface{}) {
	if myLogger.logger != nil {
		filename, funcName, line := getLogBasicInfo()
		dataStr := fmt.Sprintf(format, a...)

		myLogger.logger.Warning(fmt.Sprintf("%v %v:%v data=[%v]", filename, funcName, line, dataStr))
	}
}

func Error(format string, a ... interface{}) {
	if myLogger.logger != nil {
		filename, funcName, line := getLogBasicInfo()
		dataStr := fmt.Sprintf(format, a...)

		myLogger.logger.Error(fmt.Sprintf("%v %v:%v data=[%v]", filename, funcName, line, dataStr))
	}
}

func Fatal(format string, a ... interface{}) {
	if myLogger.logger != nil {
		filename, funcName, line := getLogBasicInfo()
		dataStr := fmt.Sprintf(format, a...)

		myLogger.logger.Critical(fmt.Sprintf("%v %v:%v data=[%v]", filename, funcName, line, dataStr))
	}
}

func LogContact(format string, a ... interface{}) {
	if myLogger.logger != nil {
		filename, funcName, line := getLogBasicInfo()
		dataStr := fmt.Sprintf(format, a...)

		myContactLogger.logger.Critical(fmt.Sprintf("%v %v:%v data=[%v]", filename, funcName, line, dataStr))
	}
}
