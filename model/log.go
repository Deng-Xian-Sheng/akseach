package model

import (
	"io"
	"log"
	"os"
)

// Logger 定义日志接口
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Close()
}

// 创建类型logger
type logger struct {
	logger *log.Logger
}

type LogMsg struct {
	MsgCurrency string `json:"msg_currency"` //通用描述
	Result      bool   `json:"result"`       //目录扫描结果
	TargetUrl   string `json:"target_url"`   //目标url
	Dir         string `json:"dir"`          //目录
}

// LogModel 日志模型
type LogModel []struct {
	Type   string `json:"type"`   //日志类型
	Time   int64  `json:"time"`   //日志时间戳
	Source string `json:"source"` //日志来源
	Msg    LogMsg `json:"msg"`    //日志信息
}

// NewLogger 创建日志实例
func NewLogger(w io.Writer) Logger {
	return &logger{
		logger: log.New(w, "", log.LstdFlags),
	}
}

// Debug 输出debug日志
func (l logger) Debug(v ...interface{}) {
	if l.logger.Writer() == os.Stdout {
		l.logger.Println(v...)
	} else {
		l.logger.Println(v...)
	}
}

// Info 输出info日志
func (l logger) Info(v ...interface{}) {
	if l.logger.Writer() == os.Stdout {
		l.logger.Println(v...)
	} else {
		l.logger.Println(v...)
	}
}

// Warn 输出warn日志
func (l logger) Warn(v ...interface{}) {
	if l.logger.Writer() == os.Stdout {
		l.logger.Println(v...)
	} else {
		l.logger.Println(v...)
	}
}

// Error 输出error日志
func (l logger) Error(v ...interface{}) {
	if l.logger.Writer() == os.Stdout {
		l.logger.Println(v...)
	} else {
		l.logger.Println(v...)
	}
}

// Fatal 输出fatal日志
func (l logger) Fatal(v ...interface{}) {
	if l.logger.Writer() == os.Stdout {
		l.logger.Println(v...)
	} else {
		l.logger.Println(v...)
	}
}

// Close 关闭日志
func (l logger) Close() {
	if l.logger.Writer() == os.Stdout {
		l.logger.Println(v...)
	} else {
		l.logger.Println(v...)
	}
}

var (
	LogOsFile, LogOsFileErr = GetOsFile("./Log.log")
	PanicLog                = log.New(io.MultiWriter(LogOsFile, os.Stdin), "[Panic]", log.Llongfile|log.LstdFlags)
	ErrorLog                = log.New(io.MultiWriter(LogOsFile, os.Stdin), "[Error]", log.Llongfile|log.LstdFlags)
	InfoLog                 = log.New(io.MultiWriter(LogOsFile, os.Stdin), "[Info]", log.Lshortfile|log.LstdFlags)
)
