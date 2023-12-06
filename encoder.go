package log

import (
	"github.com/allen-shaw/log/Internal/encoder"
	"go.uber.org/zap/zapcore"
)

const (
	_timeLayout = "2006-01-02 15:04:05.999-07:00"
	_separator  = "|"
)

func newProductionEncoderConfig(traceKey string) encoder.Config {
	return encoder.Config{
		TraceKey: traceKey,
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

func newEncoder(opt *options) zapcore.Encoder {
	ecfg := newProductionEncoderConfig(opt.traceKey)
	return newConsoleEncoder(ecfg)
}

func newConsoleEncoder(cfg encoder.Config) zapcore.Encoder {
	return encoder.NewConsoleEncoder(cfg)
}
