package logger

import (
	"github.com/sirupsen/logrus"
)

func LogErr(param logrus.Fields, err error) {
	const levelStackTrace = 3
	frame := TraceLevel(levelStackTrace)

	param["func"] = frame.Function
	param["file"] = frame.File
	param["line"] = frame.Line

	logrus.WithFields(param).Error(err)
}

func LogWarn(param logrus.Fields, msg string) {
	logrus.WithFields(param).Warn(msg)
}
