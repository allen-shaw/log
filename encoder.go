package log

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const (
	_traceKey   = "@trace_id"
	_timeLayout = "2006-01-02 15:04:05.999-07:00"
	_separator  = "|"
)

type encoderConfig struct {
	TraceKey string `json:"trace_key"`
	zapcore.EncoderConfig
}

func newProductionEncoderConfig() encoderConfig {
	return encoderConfig{
		TraceKey: _traceKey,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:          "ts",
			LevelKey:         "level",
			NameKey:          "logger",
			CallerKey:        "caller",
			FunctionKey:      "func",
			MessageKey:       "msg",
			StacktraceKey:    "stacktrace",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      zapcore.LowercaseColorLevelEncoder,
			EncodeTime:       zapcore.TimeEncoderOfLayout(_timeLayout),
			EncodeDuration:   zapcore.MillisDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			ConsoleSeparator: _separator,
		},
	}
}

type consoleEncoder struct {
	*encoderConfig
	buf *buffer.Buffer
}

func newEncoder(opt *options) zapcore.Encoder {
	ecfg := newProductionEncoderConfig()
	return newConsoleEncoder(ecfg)
}

func newConsoleEncoder(cfg encoderConfig) zapcore.Encoder {
	panic("no implement")
}
