package log

import "go.uber.org/zap/zapcore"

type Level = zapcore.Level

const (
	unknownLevel = zapcore.DebugLevel - 1
	DebugLevel   = zapcore.DebugLevel
	InfoLevel    = zapcore.InfoLevel
	WarnLevel    = zapcore.WarnLevel
	ErrorLevel   = zapcore.ErrorLevel
	DPanicLevel  = zapcore.DPanicLevel
	PanicLevel   = zapcore.PanicLevel
	FatalLevel   = zapcore.FatalLevel
)
