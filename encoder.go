package log

import (
	"github.com/allen-shaw/log/Internal/encoder"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const (
	_traceKey   = "@trace_id"
	_timeLayout = "2006-01-02 15:04:05.999-07:00"
	_separator  = "|"
)

func newProductionEncoderConfig() encoder.Config {
	return encoder.Config{
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
	*encoder.Config
	buf *buffer.Buffer
}

func newEncoder(opt *options) zapcore.Encoder {
	ecfg := newProductionEncoderConfig()
	return newConsoleEncoder(ecfg)
}

func newConsoleEncoder(cfg encoder.Config) zapcore.Encoder {
	return encoder.NewConsoleEncoder(cfg)
}
