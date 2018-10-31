package common

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"os"
	"syscall"
	"time"
)

var (
	Logger  echo.Logger
	stdFile *os.File
	logPath string
)

// LoggerInit init globel logger
func LoggerInit(e *echo.Echo, logpath string, debug bool) {
	Logger = e.Logger
	logPath = logpath
	if debug {
		Logger.SetLevel(log.DEBUG)
	} else {
		Logger.SetLevel(log.INFO)
	}

	//写入文件
	UpLogFile()

}

func UpLogFile() {
	setLogFile()
	Logger.SetOutput(stdFile)
}

func setLogFile() {
	strtime := time.Now().Format("2006-01-02") //时间模板 2006-01-02.15.04.05  目前按照天来记日志
	filename := logPath + "." + strtime + ".log"
	filetemp := stdFile
	var err error
	stdFile, err = os.OpenFile(filename, syscall.O_CREAT|syscall.O_APPEND|syscall.O_RDWR, 0666) //打开文件
	if err != nil {
		Logger.Error("set log file error: ", err)
		return
	}

	if filetemp != nil {
		filetemp.Close()
	}
}
