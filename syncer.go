package log

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

func newFileSyncer(tee *teeOption, localTime bool) io.Writer {
	syncer := &lumberjack.Logger{
		Filename:   tee.file,
		MaxSize:    tee.ratate.maxSize,
		MaxAge:     tee.ratate.maxAge,
		MaxBackups: tee.ratate.maxBackups,
		LocalTime:  localTime,
		Compress:   tee.ratate.compress,
	}
	return syncer
}
