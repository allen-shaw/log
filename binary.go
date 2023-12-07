package log

import (
	"encoding/base64"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func binaries(key string, ba [][]byte) zap.Field {
	return zap.Array(key, binaryArray(ba))
}

type binaryArray [][]byte

func (b binaryArray) MarshalLogArray(en zapcore.ArrayEncoder) error {
	for _, byt := range b {
		en.AppendString(base64.StdEncoding.EncodeToString(byt))
	}

	return nil
}
