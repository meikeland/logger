package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func getFormatter() logrus.Formatter {
	return &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			fileName := fmt.Sprintf("%s:%d", f.File, f.Line)
			return funcname, fileName
		},
	}
}

func newLogrusLogger(config Config) (*logrus.Logger, error) {
	logLevel := config.Level
	if logLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	// 控制台输出
	stdOutHandler := os.Stdout

	lLogger := &logrus.Logger{
		Out:       stdOutHandler, // 默认将日志打印到控制台，无论设置与否
		Formatter: getFormatter(),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	// 日志文件输出
	if config.EnableFile {
		src, err := os.OpenFile(config.FileLocation, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		if config.EnableConsole {
			lLogger.SetOutput(io.MultiWriter(stdOutHandler, src))
		} else {
			lLogger.SetOutput(src)
		}
	}

	if config.AppendCaller {
		lLogger.SetReportCaller(true)
	}

	return lLogger, nil
}
