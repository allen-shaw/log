package encoder

import "go.uber.org/zap/zapcore"

type Config struct {
	TraceKey string `json:"traceKey" yaml:"traceKey"`
	zapcore.EncoderConfig
}
