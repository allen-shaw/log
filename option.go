package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option zap.Option

type options struct {
	level     Level
	debug     bool
	localTime bool
	tees      []teeOption
}

type teeOption struct {
	file   string // path+filename
	level  Level
	skip   int
	hooks  []func(zapcore.Entry) error
	opts   []options
	ratate rotateOption
}

type rotateOption struct {
	compress   bool
	maxSize    int
	maxAge     int
	maxBackups int
}

var (
	defRotateOpts = rotateOption{
		compress:   false,
		maxSize:    100,
		maxAge:     7,
		maxBackups: 10,
	}
	defTeeOpts = teeOption{
		file:   "./path/server.log",
		level:  unknownLevel,
		ratate: defRotateOpts,
	}
	defOpts = options{
		level:     InfoLevel,
		localTime: true,
		tees:      []teeOption{defTeeOpts},
	}
)

func evalOptions(opts ...Options) *options {
	optCopy := defOpts
	for _, o := range opts {
		o.apply(&optCopy)
	}
	return &optCopy
}

type Options interface {
	apply(*options)
}

func WithLevel(l Level) Options {
	return levelOption(l)
}

func WithDebug(debug bool) Options {
	return debugOption(debug)
}

func WithLocalTime(localTime bool) Options {
	return localTimeOption(localTime)
}

// internal

type levelOption Level

func (l levelOption) apply(opts *options) {
	opts.level = Level(l)
}

type debugOption bool

func (d debugOption) apply(opts *options) {
	opts.debug = bool(d)
}

type localTimeOption bool

func (lt localTimeOption) apply(opts *options) {
	opts.localTime = bool(lt)
}
