package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

//type Logger struct {
//	internalLogger *logrus.Logger
//}
//
//func New() *Logger{
//	return &Logger{logrus.New()}
//}
//
//func (this *Logger) Debug(args ...interface{}){
//	this.internalLogger.Debug(args)
//}
//
//func (this *Logger) Error(args ...interface{}){
//	this.internalLogger.Error(args)
//}

type Level uint32

const(
	Identity = "Yggdrasil"
	ErrorLevel = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

var appname string = "unknow"
var version int = 1
var subversion int = 1
var pid int = 0
var hostname string = ""

func Setup(app string, ver int, subver int){
	appname = app
	version = ver
	subversion = subver
	pid = os.Getpid()
	hostname, _ = os.Hostname()

	logrus.SetFormatter(&CommonFormatter{&logrus.JSONFormatter{}})
}

func SetLevel(level Level){
	logrus.SetLevel(convertLevel(level))
}

func convertLevel(level Level) logrus.Level{
	if level == ErrorLevel{
		return logrus.ErrorLevel
	}else if level == WarnLevel{
		return logrus.WarnLevel
	}else if level == InfoLevel{
		return logrus.InfoLevel
	}else {
		return logrus.DebugLevel
	}
}


func Error(args ...interface{}){

}

type JSONFormatter struct {
}
