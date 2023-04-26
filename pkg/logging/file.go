package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/" //日志文件的保存路径
	LogSaveName = "log"           //日志文件的名字
	LogFileExt  = "log"           //日志文件拓展名
	TimeFormat  = "20060102"      //日志文件时间格式
)

// 获取日志文件的保存路径
func getLogFilePath() string {
	return fmt.Sprint("%s", LogSavePath)
}

// 获取日志文件的完整路径
func getLogFileFullPath() string {
	prefixPath := getLogFilePath() //日志前缀：保存路径
	suffixPath := fmt.Sprint("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	//拼接为：名字+时间.拓展名
	return fmt.Sprint("%s%s", prefixPath, suffixPath)
}

// 打开并且返回一个日志文件的句柄，在其他日志操作中使用该句柄，以便于继续向日志文件中追加内容。
func openLogFile(filePath string) *os.File {
	//stat检查是否存在
	_, err := os.Stat(filePath) //错误信息包括路径，错误等信息
	//下面是判断错误类型
	switch {
	case os.IsNotExist(err): //判断给定的错误是否表示文件或目录不存在，返回一个布尔值
		mkDir() //不存在就创建文件保存目录
	case os.IsPermission(err): //判断给定的错误是否表示权限不足，返回一个布尔值
		log.Fatal("Permission :%v", err) //是权限不足，就打印日志错误log.fatal
	}
	//打开文件，保存斌且给返回句柄handle
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//openFile打开文件    设置打开模式为追加写入  ，如果不存在就创建，                文件权限为0644
	if err != nil {
		log.Fatal("Fail to OpenFile :%v", err)
	}
	//返回文件句柄
	return handle
}

// 创建日志文件的保存目录
func mkDir() {
	dir, _ := os.Getwd() //获取当前目录对应根路径名 get working directory
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	//mkdir all,用于逐级创建目录，创建对应的目录以及所需的子目录，若成功则返回nil，否则返回error
	//拼接：工作目录+/+日志文件保存目录
	//ModePerm  permission mode  设置该目录的权限为所有者可以读写的权限(0777)。
	if err != nil {
		panic(err)
	}
}
