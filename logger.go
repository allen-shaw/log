package log

import (
	"sync/atomic"

	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
	level  atomic.Int32 // Level type
}

func NewLogger(opts ...Options) *Logger {
	opt := evalOptions(opts...)
	l := &Logger{}
	l.level.Store(int32(opt.level))

	// encoder := newEncoder(opt)

	return l
}
