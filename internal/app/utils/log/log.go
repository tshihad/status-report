package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

//Defaultloggger for logs
var Defaultloggger = logrus.New()

func ServiceLogger(level logrus.Level, name string, serviceDetails interface{}) *logrus.Entry {
	logger := logrus.New()
	f, err := os.OpenFile("logs/logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic("Cant open logs/logs.txt " + err.Error())
	}
	logger.Out = f
	logger.Level = level
	return logger.WithFields(logrus.Fields{
		"name":            name,
		"service_details": serviceDetails,
	})
}

func Infof(format string, args ...string) {
	Defaultloggger.Infof(format, args)
}

func Info(args ...string) {
	Defaultloggger.Info(args)
}

func Debugf(format string, args ...string) {
	Defaultloggger.Debugf(format, args)
}

func Debug(args ...string) {
	Defaultloggger.Debug(args)
}

func Warnf(format string, args ...string) {
	Defaultloggger.Warnf(format, args)
}

func Warn(args ...string) {
	Defaultloggger.Warn(args)
}

func Errorf(format string, args ...string) {
	Defaultloggger.Errorf(format, args)
}

func Error(args ...string) {
	Defaultloggger.Error(args)
}

func Print(args ...string) {
	Defaultloggger.Print(args)
}

func Printf(format string, args ...string) {
	Defaultloggger.Printf(format, args)
}

func Panicf(format string, args ...string) {
	Defaultloggger.Panicf(format, args)
}

func Panic(args ...string) {
	Defaultloggger.Panic(args)
}
