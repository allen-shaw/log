package log

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type tee struct {
	level   zap.LevelEnablerFunc
	syncer  io.Writer
	encoder zapcore.Encoder
}

func newTee(encoder zapcore.Encoder, opts *options) []*tee {
	var tops []*tee
	for _, t := range opts.tees {
		syncer := newFileSyncer(&t, opts.localTime)
		top := &tee{
			level:   t.levelFunc(opts.level),
			syncer:  syncer,
			encoder: encoder,
		}
		tops = append(tops, top)
	}
	return tops
}
