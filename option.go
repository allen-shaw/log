package log

import (
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const _logSuffix = ".log"

type Option zap.Option

type options struct {
	level     Level
	debug     bool
	localTime bool
	tees      []teeOption
	skip      int
	hooks     []func(zapcore.Entry) error
	opts      []zap.Option
}

type LevelEnablerFunc func(Level) zap.LevelEnablerFunc

type teeOption struct {
	file      string // path+filename
	levelFunc LevelEnablerFunc
	ratate    rotateOption
}

type rotateOption struct {
	compress   bool
	maxSize    int
	maxAge     int
	maxBackups int
}

var (
	defLevelFunc = func(level Level) zap.LevelEnablerFunc {
		return func(l zapcore.Level) bool {
			return l >= level
		}
	}
	defRotateOpts = rotateOption{
		compress:   false,
		maxSize:    100,
		maxAge:     7,
		maxBackups: 10,
	}
	defTeeOpts = teeOption{
		file:      "./log/server.log",
		levelFunc: defLevelFunc,
		ratate:    defRotateOpts,
	}
	defOpts = options{
		level:     InfoLevel,
		localTime: true,
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

func WithSkip(skip int) Options {
	return skipOption(skip)
}

func WithHooks(hooks []func(zapcore.Entry) error) Options {
	return hooksOption(hooks)
}

func WithOptions(opts []zap.Option) Options {
	return optOptions(opts)
}

type fileOption interface {
	apply(*rotateOption)
}

func WithFile(file string, f LevelEnablerFunc, opts ...fileOption) Options {
	if path.Ext(file) != _logSuffix {
		file += _logSuffix
	}

	rotp := defRotateOpts
	for _, o := range opts {
		o.apply(&rotp)
	}

	tee := teeOption{
		file:      file,
		levelFunc: f,
		ratate:    rotp,
	}
	return teeOptions(tee)
}

func WithMaxSize(maxSize int) fileOption {
	return maxSizeOptions(maxSize)
}

func WithMaxAge(age int) fileOption {
	return maxAgeOptions(age)
}

func WithMaxBackups(backup int) fileOption {
	return maxBackupsOptions(backup)
}

func WithCompress(compress bool) fileOption {
	return compressOptions(compress)
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

type skipOption int

func (s skipOption) apply(opts *options) {
	opts.skip = int(s)
}

type hooksOption []func(zapcore.Entry) error

func (h hooksOption) apply(opts *options) {
	opts.hooks = []func(zapcore.Entry) error(h)
}

type optOptions []zap.Option

func (o optOptions) apply(opts *options) {
	opts.opts = optOptions(o)
}

type teeOptions teeOption

func (t teeOptions) apply(opts *options) {
	opts.tees = append(opts.tees, teeOption(t))
}

type maxSizeOptions int

func (m maxSizeOptions) apply(opts *rotateOption) {
	opts.maxSize = int(m)
}

type maxAgeOptions int

func (m maxAgeOptions) apply(opts *rotateOption) {
	opts.maxAge = int(m)
}

type maxBackupsOptions int

func (m maxBackupsOptions) apply(opts *rotateOption) {
	opts.maxBackups = int(m)
}

type compressOptions bool

func (c compressOptions) apply(opts *rotateOption) {
	opts.compress = bool(c)
}
