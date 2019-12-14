package logger

import (
	"github.com/sirupsen/logrus"
)

// 全局的log对象
var log *logrus.Logger

// Fields 以withfield形式打印的日志项
type Fields map[string]interface{}

// Config 控制日志组件的打印信息
type Config struct {
	EnableConsole bool   // 是否在控制台打印日志
	EnableFile    bool   // 是否在文件记录日志
	Level         string // 文件日志级别
	FileLocation  string // 文件路径
	AppendCaller  bool   // 是否打印代码位置
}

//New 创建一个新的日志组件
func New(config Config) error {
	logger, err := newLogrusLogger(config)
	if err != nil {
		return err
	}

	log = logger
	return nil
}

// Debugf 打印调试日志
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof 打印信息日志
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf 打印警告日志
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf 打印错误日志
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf 打印严重错误日志
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Panicf 打印错误日志，并结束运行
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// WithFields 以字段形式输出
func WithFields(keyValues Fields) *logrus.Entry {
	return log.WithFields(convertToLogrusFields(keyValues))
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
