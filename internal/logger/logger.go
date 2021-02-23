package logger

import (
	"os"
	"runtime"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

// Лог будет вывотить и в stdout и в файл /var/log/facechain/pipelinego.log
// Настроен logrotate

func InitLog(loglevel string, nameApp string, folderLog string) { //nolint:funlen
	var logLevel = logrus.InfoLevel //nolint:ineffassign

	switch loglevel {
	case "info":
		logLevel = logrus.InfoLevel
	case "trace":
		logLevel = logrus.TraceLevel
	case "debug":
		logLevel = logrus.DebugLevel
	case "warning":
		logLevel = logrus.WarnLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	case "fatal":
		logLevel = logrus.FatalLevel
	default:
		logLevel = logrus.InfoLevel

		frame := Trace()

		logrus.WithFields(logrus.Fields{
			"func": frame.Function,
			"file": frame.File,
			"line": frame.Line,
		}).Error("Wrong log level")
	}

	if _, err := os.Stat(folderLog); os.IsNotExist(err) {
		os.Mkdir(folderLog, 0766)

		frame := Trace()

		logrus.WithFields(logrus.Fields{
			"err_message": err,
			"file":        frame.File,
			"line":        frame.Line,
			"func":        frame.Function,
		}).Warn("No found folder ", folderLog)
	}

	const maxSizeLog = 50

	const maxBackups = 3

	const maxAge = 28

	// Настраиваем logrotate
	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   folderLog + "/" + nameApp + ".log",
		MaxSize:    maxSizeLog, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, //days
		Level:      logLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		frame := Trace()

		logrus.WithFields(logrus.Fields{
			"err_message": err,
			"file":        frame.File,
			"line":        frame.Line,
			"func":        frame.Function,
		}).Fatal("Failed to initialize file rotate hook: ")
	}

	// Отельно настроим вывод в stdout
	logrus.SetLevel(logLevel)
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	logrus.AddHook(rotateFileHook)

	frame := Trace()

	logrus.WithFields(logrus.Fields{
		"func": frame.Function,
	}).Info("Init log")
}

// Дает информацию о строке,
// файле и функции где произошла ошибка

func Trace() runtime.Frame {
	const LevelStrace = 2

	pc := make([]uintptr, 15)

	n := runtime.Callers(LevelStrace, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return frame
}

// Дает информацию о строке, файле
// и функции где произошла ошибка

func TraceLevel(l int) runtime.Frame {
	pc := make([]uintptr, 15)
	n := runtime.Callers(l, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return frame
}

// example
// Не использовать эту функции потому что она указывает на саму себя
//func printTrace() {
//frame := Trace()
//logrus.WithFields(logrus.Fields{
//"file": frame.File,
//"line": frame.Line,
//"func": frame.Function,
//}).Trace("")
//}
