package log

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	internalLogger *logrus.Logger
}

func (this *Logger) Debug(args ...interface{}){
	this.internalLogger.Debug(args)
}

func (this *Logger) Debugf(format string, args ...interface{}){
	this.internalLogger.Debugf(format, args)
}

func (this *Logger) Info(args ...interface{}){
	this.internalLogger.Info(args)
}

func (this *Logger) Infof(format string, args ...interface{}){
	this.internalLogger.Infof(format, args)
}

func (this *Logger) Warn(args ...interface{}){
	this.internalLogger.Warn(args)
}

func (this *Logger) Warnf(format string, args ...interface{}){
	this.internalLogger.Warnf(format, args)
}

func (this *Logger) Error(args ...interface{}){
	this.internalLogger.Error(args)
}

func (this *Logger) Errorf(format string, args ...interface{}){
	this.internalLogger.Errorf(format, args)
}


