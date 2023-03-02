package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

/**
* @Author $
* @Description //TODO $
* @Date $ $
* @Param $
* @return $
**/

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// String
//
//	@Description: 错误等级字典
//	@receiver l
//	@return string
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

// NewLogger
//
//	@Description: 初始化logger
//	@Params $w
//	@params $prefix
//	@params $flag
//	@return *Logger
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

// clone
//
//	@Description: 深拷贝对象
//	@receiver l
//	@return *Logger
func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// WithFields
//
//	@Description: 设置日志公共字段
//	@receiver l
//	@params $f
//	@return *Logger
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}

	for k, v := range f {
		ll.fields[k] = v
	}

	return ll
}

// WithContext
//
//	@Description: 设置日志上下文
//	@receiver l
//	@params $ctx
//	@return *Logger
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// WithCaller
//
//	@Description: 设置当前某一层调用栈的信息（程序计数器、文件信息、行号）
//	@receiver l
//	@params $skip
//	@return *Logger
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return ll
}

// WithCallersFrames
//
//	@Description: 设置当前的整个调用栈信息
//	@receiver l
//	@return *Logger
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}

	ll := l.clone()
	ll.callers = callers
	return ll
}

// JSONFormat
//
//	@Description: 编写日志内容的格式化
//	@receiver l
//	@params $level
//	@params $message
//	@return map[string]interface{}
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers

	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

// 链路追踪
func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Trace-ID"),
		})
	}
	return l
}

// Output
//
//	@Description: 日志输出
//	@receiver l
//	@params $level
//	@params $message
func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

// Info
//
//	@Description: info日志输出
//	@receiver l
//	@params $v
func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelInfo, fmt.Sprint(v...))
}

// Infof
//
//	@Description: info日志格式化输出
//	@receiver l
//	@params $format
//	@params $v
func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Output(LevelPanic, fmt.Sprintf(format, v...))
}
