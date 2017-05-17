package log

import (
	"os"
	"github.com/sirupsen/logrus"
)

type Level uint32

const(
	Identity = "Yggdrasil"
	ErrorLevel = iota
	InfoLevel
	WarnLevel
	DebugLevel
)

var appname string = "unknow"
var version int = 1
var subversion int = 1
var pid int = 0
var hostname string = ""
var rootLevel Level = InfoLevel

func Setup(app string, ver int, subver int){
	appname = app
	version = ver
	subversion = subver
	pid = os.Getpid()
	hostname, _ = os.Hostname()
}

func SetRootLevel(level Level){
	rootLevel = level
}

func New() *Logger{
	logger := logrus.New()
	logger.Level = convertLevel(rootLevel)
	logger.Formatter = &CommonFormatter{logrus.JSONFormatter{}}

	return &Logger{logger}
}

func convertLevel(level Level) logrus.Level{
	if level == ErrorLevel{
		return logrus.ErrorLevel
	}else if level == WarnLevel{
		return logrus.WarnLevel
	}else if level == InfoLevel {
		return logrus.InfoLevel
	}

	return logrus.DebugLevel
}

