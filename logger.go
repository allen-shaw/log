package log

import (
	"sync/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
	level  atomic.Int32 // Level type
}

func NewLogger(opts ...Options) *Logger {
	opt := evalOptions(opts...)

	l := &Logger{}
	l.level.Store(int32(opt.level))

	if len(opt.tees) == 0 {
		opt.tees = append(opt.tees, defTeeOpts)
	}

	encoder := newEncoder(opt)
	tops := newTee(encoder, opt)
	lg := newLogger(tops, opt.skip, opt.hooks, opt.opts...)
	l.logger = lg

	return l
}

func newLogger(tops []*tee, skip int, hooks []func(zapcore.Entry) error, opts ...zap.Option) *zap.Logger {
	var cores []zapcore.Core

	for _, top := range tops {
		core := zapcore.NewCore(top.encoder, zapcore.AddSync(top.syncer), top.level)
		cores = append(cores, core)
	}

	opt := append([]zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(PanicLevel),
		zap.AddCallerSkip(skip),
		zap.Hooks(hooks...),
	}, opts...)

	logger := zap.New(zapcore.NewTee(cores...), opt...)
	logger = logger.With(zap.String(_traceKey, "-"))

	return logger
}

func (l *Logger) GetLevel() Level {
	return Level(l.level.Load())
}

func (l *Logger) SetLevel(lvl Level) {
	l.level.Store(int32(lvl))
}

func (l *Logger) SetFields(fields ...Field) {
	lg := l.logger.With(fields...)
	l.logger = lg
}

func (l *Logger) SetOptions(opts ...zap.Option) {
	lg := l.logger.WithOptions(opts...)
	l.logger = lg
}

func (l *Logger) With(fields ...Field) *Logger {
	lg := l.logger.With(fields...)
	log := l.clone()
	log.logger = lg
	return log
}

func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	lg := l.logger.WithOptions(opts...)
	log := l.clone()
	log.logger = lg
	return log
}

func (l *Logger) Logger() *zap.Logger {
	return l.logger
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...Field) {
	l.logger.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.logger.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}

func (l *Logger) clone() *Logger {
	log := &Logger{}
	log.logger = l.logger
	log.level.Store(l.level.Load())
	return log
}

func (l *Logger) Sugar() *zap.SugaredLogger {
	return l.logger.Sugar()
}
