package logs

import "github.com/sirupsen/logrus"

// Default default log
var Default = NewLogger()

// Logger interface
type Logger interface {
	logrus.FieldLogger
}

func Info(args ...interface{}) {
	Default.Info(args...)
}

func Infoln(args ...interface{}) {
	Default.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	Default.Infof(format, args...)
}

func Debug(args ...interface{}) {
	Default.Debug(args...)
}

func Debugln(args ...interface{}) {
	Default.Debugln(args...)
}

func Debugf(format string, args ...interface{}) {
	Default.Debugf(format, args...)
}

func Error(args ...interface{}) {
	Default.Error(args)
}

func Errorln(args ...interface{}) {
	Default.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	Default.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	Default.Fatal(args)
}

func Fatalln(args ...interface{}) {
	Default.Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	Default.Fatalf(format, args...)
}
