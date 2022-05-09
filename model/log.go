// Package model.Logger
// 用于日志分级、美化控制台输出、持久化存储
package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

// Logger 定义日志接口
type Logger interface {
	// Debug 输出调试日志
	Debug(v ...interface{})
	// Info 输出信息日志
	Info(v ...interface{})
	// Warn 输出警告日志
	Warn(v ...interface{})
	// Error 输出错误日志
	Error(v ...interface{})
	// Fatal 输出致命错误日志
	Fatal(v ...interface{})
	// SetLevel 设置日志级别
	// level: 0-debug, 1-info, 2-warn, 3-error, 4-fatal, 5-off
	SetLevel(level int)
	// Close 关闭日志
	Close()
}

const (
	DEBUG = iota // 调试日志
	INFO         // 信息日志
	WARN         // 警告日志
	ERROR        // 错误日志
	FATAL        // 致命错误日志
	NOLOG        // 关闭日志
)

// logger 定义日志结构体
type logger struct {
	logger      *log.Logger
	persistence *[]io.Writer // 持久化输出
	cmd         *[]io.Writer // 控制台输出
	level       int          // 日志级别
}

// logModel 日志模型
type logModel struct {
	Type       string `json:"type"`        //日志类型
	Time       int64  `json:"time"`        //日志时间戳
	SourceFile string `json:"source"`      //日志来源的文件名
	SourceLine int    `json:"source_line"` //日志来源的行数
	SourceFunc string `json:"source_func"` //日志来源的函数
	Msg        string `json:"msg"`         //日志信息
}

// 结构化
func structured(Type string, v []interface{}, pc uintptr, file string, line int, ok bool) *logModel {
	var logmodel logModel
	logmodel.Type = Type
	logmodel.Time = time.Now().UnixMicro()
	if ok {
		logmodel.SourceFile = file
		logmodel.SourceLine = line
		logmodel.SourceFunc = runtime.FuncForPC(pc).Name()
	}
	logmodel.Msg = fmt.Sprint(v...)
	return &logmodel
}

// 日志美化
func colored(model *logModel) string {
	var msg, color, removeColor string
	switch model.Type {
	case "DEBUG":
		// 判断运行环境
		if runtime.GOOS != "windows" {
			color = "\u001B[1;36m"
			removeColor = "\u001B[0m"
		}
		msg = fmt.Sprintf("%s[%s]%s %s %s|File:%s %s %s|Line:%s %d %s|Func:%s %s %s|Msg:%s %s", color, model.Type, removeColor, time.UnixMicro(model.Time).Format("2006-01-02 15:04:05.999"), color, removeColor, model.SourceFile, color, removeColor, model.SourceLine, color, removeColor, model.SourceFunc, color, removeColor, model.Msg)
	case "INFO":
		if runtime.GOOS != "windows" {
			color = "\u001B[1;32m"
			removeColor = "\u001B[0m"
		}
		msg = fmt.Sprintf("%s[%s]%s %s %s|File:%s %s %s|Line:%s %d %s|Func:%s %s %s|Msg:%s %s", color, model.Type, removeColor, time.UnixMicro(model.Time).Format("2006-01-02 15:04:05.999"), color, removeColor, model.SourceFile, color, removeColor, model.SourceLine, color, removeColor, model.SourceFunc, color, removeColor, model.Msg)
	case "WARN":
		if runtime.GOOS != "windows" {
			color = "\u001B[1;33m"
			removeColor = "\u001B[0m"
		}
		msg = fmt.Sprintf("%s[%s]%s %s %s|File:%s %s %s|Line:%s %d %s|Func:%s %s %s|Msg:%s %s", color, model.Type, removeColor, time.UnixMicro(model.Time).Format("2006-01-02 15:04:05.999"), color, removeColor, model.SourceFile, color, removeColor, model.SourceLine, color, removeColor, model.SourceFunc, color, removeColor, model.Msg)
	case "ERROR":
		if runtime.GOOS != "windows" {
			color = "\u001B[1;31m"
			removeColor = "\u001B[0m"
		}
		msg = fmt.Sprintf("%s[%s]%s %s %s|File:%s %s %s|Line:%s %d %s|Func:%s %s %s|Msg:%s %s", color, model.Type, removeColor, time.UnixMicro(model.Time).Format("2006-01-02 15:04:05.999"), color, removeColor, model.SourceFile, color, removeColor, model.SourceLine, color, removeColor, model.SourceFunc, color, removeColor, model.Msg)
	case "FATAL":
		if runtime.GOOS != "windows" {
			color = "\u001B[1;35m"
			removeColor = "\u001B[0m"
		}
		msg = fmt.Sprintf("%s[%s]%s %s %s|File:%s %s %s|Line:%s %d %s|Func:%s %s %s|Msg:%s %s", color, model.Type, removeColor, time.UnixMicro(model.Time).Format("2006-01-02 15:04:05.999"), color, removeColor, model.SourceFile, color, removeColor, model.SourceLine, color, removeColor, model.SourceFunc, color, removeColor, model.Msg)
	default:
		if runtime.GOOS != "windows" {
			color = "\u001B[1;37m"
			removeColor = "\u001B[0m"
		}
		msg = fmt.Sprintf("%s[%s]%s %s %s|File:%s %s %s|Line:%s %d %s|Func:%s %s %s|Msg:%s %s", color, model.Type, removeColor, time.UnixMicro(model.Time).Format("2006-01-02 15:04:05.999"), color, removeColor, model.SourceFile, color, removeColor, model.SourceLine, color, removeColor, model.SourceFunc, color, removeColor, model.Msg)
	}
	return msg
}

// NewLogger 创建日志实例
// cmd 控制台输出
// persistence 持久化输出
func NewLogger(cmd, persistence []io.Writer) Logger {
	return &logger{
		logger:      &log.Logger{},
		persistence: &persistence,
		cmd:         &cmd,
		level:       FATAL, // 默认级别
	}
}

// Debug 输出debug日志
func (l *logger) Debug(v ...interface{}) {
	// 判断日志级别
	if l.level <= DEBUG {
		// 获取调用者的信息
		pc, file, line, ok := runtime.Caller(1)
		// 结构化
		model := structured("DEBUG", v, pc, file, line, ok)
		// 判断是否持久化
		if len(*l.persistence) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.persistence...))
			marshal, err := json.Marshal(&model)
			if err != nil {
				return
			}
			err = l.logger.Output(2, string(marshal))
			if err != nil {
				return
			}
		}
		// 判断是否控制台输出
		if len(*l.cmd) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.cmd...))
			s := colored(model)
			err := l.logger.Output(2, s)
			if err != nil {
				return
			}
		}
	}
}

// Info 输出info日志
func (l *logger) Info(v ...interface{}) {
	// 判断日志级别
	if l.level <= INFO {
		// 获取调用者的信息
		pc, file, line, ok := runtime.Caller(1)
		// 结构化
		model := structured("INFO", v, pc, file, line, ok)
		// 判断是否持久化
		if len(*l.persistence) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.persistence...))
			marshal, err := json.Marshal(&model)
			if err != nil {
				return
			}
			err = l.logger.Output(2, string(marshal))
			if err != nil {
				return
			}
		}
		// 判断是否控制台输出
		if len(*l.cmd) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.cmd...))
			s := colored(model)
			err := l.logger.Output(2, s)
			if err != nil {
				return
			}
		}
	}
}

// Warn 输出warn日志
func (l *logger) Warn(v ...interface{}) {
	// 判断日志级别
	if l.level <= WARN {
		// 获取调用者的信息
		pc, file, line, ok := runtime.Caller(1)
		// 结构化
		model := structured("WARN", v, pc, file, line, ok)
		// 判断是否持久化
		if len(*l.persistence) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.persistence...))
			marshal, err := json.Marshal(&model)
			if err != nil {
				return
			}
			err = l.logger.Output(2, string(marshal))
			if err != nil {
				return
			}
		}
		// 判断是否控制台输出
		if len(*l.cmd) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.cmd...))
			s := colored(model)
			err := l.logger.Output(2, s)
			if err != nil {
				return
			}
		}
	}
}

// Error 输出error日志
func (l *logger) Error(v ...interface{}) {
	// 判断日志级别
	if l.level <= ERROR {
		// 获取调用者的信息
		pc, file, line, ok := runtime.Caller(1)
		// 结构化
		model := structured("ERROR", v, pc, file, line, ok)
		// 判断是否持久化
		if len(*l.persistence) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.persistence...))
			marshal, err := json.Marshal(&model)
			if err != nil {
				return
			}
			err = l.logger.Output(2, string(marshal))
			if err != nil {
				return
			}
		}
		// 判断是否控制台输出
		if len(*l.cmd) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.cmd...))
			s := colored(model)
			err := l.logger.Output(2, s)
			if err != nil {
				return
			}
		}
	}
}

// Fatal 输出fatal日志
func (l *logger) Fatal(v ...interface{}) {
	// 判断日志级别
	if l.level <= FATAL {
		// 获取调用者的信息
		pc, file, line, ok := runtime.Caller(1)
		// 结构化
		model := structured("FATAL", v, pc, file, line, ok)
		// 判断是否持久化
		if len(*l.persistence) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.persistence...))
			marshal, err := json.Marshal(&model)
			if err != nil {
				return
			}
			err = l.logger.Output(2, string(marshal))
			if err != nil {
				return
			}
		}
		// 判断是否控制台输出
		if len(*l.cmd) > 0 {
			l.logger.SetOutput(io.MultiWriter(*l.cmd...))
			s := colored(model)
			err := l.logger.Output(2, s)
			if err != nil {
				return
			}
		}
		// 调用panic而不是os.Exit(1)；因为os.Exit(1)不能被recover()处理
		panic(v...)
	}
}

// SetLevel 设置日志级别
func (l *logger) SetLevel(level int) {
	l.level = level
}

// Close 关闭日志
func (l *logger) Close() {
	l.level = NOLOG
}

var (
	LogOsFile *os.File
	LogWendy  Logger
)

func init() {
	var LogOsFileErr error
	LogOsFile, LogOsFileErr = GetOsFile("./Log.log")
	if LogOsFileErr != nil {
		panic(LogOsFileErr)
	}
	LogWendy = NewLogger([]io.Writer{os.Stdout}, []io.Writer{LogOsFile})
}
