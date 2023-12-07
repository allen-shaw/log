package examples

import (
	"testing"

	"github.com/allen-shaw/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestDefaultLogger(t *testing.T) {
	logger := log.NewLogger()
	logger.Info("hello", log.String("test_str", "world"))

	logger.Sync()
}

func TestLevelLogger(t *testing.T) {
	logger := log.NewLogger(
		log.WithLevel(log.InfoLevel),
		log.WithSkip(1),
		log.WithFile(
			"./logs/info",
			func(l log.Level) zap.LevelEnablerFunc {
				return func(l zapcore.Level) bool { return l == log.InfoLevel }
			},
			log.WithMaxBackups(5),
		),
		log.WithFile(
			"./logs/error",
			func(l log.Level) zap.LevelEnablerFunc {
				return func(l zapcore.Level) bool { return l == log.ErrorLevel }
			},
			log.WithMaxBackups(3),
		),
	)

	logger.Debug("test_debug", log.Int("debug", 1)) // 不会输出
	logger.Info("test_info", log.Int("info", 2))
	logger.Warn("test_warn", log.Int("warn", 3)) // 不会输出
	logger.Error("test_error", log.Int("error", 4))

	logger.Sync()
}

func TestWithField(t *testing.T) {
	logger := log.NewLogger().With(log.Int("id", 123))
	defer logger.Sync()

	logger.Info("hello", log.String("test_str", "world"))
}

func TestWithTraceID(t *testing.T) {
	myTraceKey := "my_trace_id"
	logger := log.NewLogger(
		log.WithSkip(1),
		log.WithTraceKey(myTraceKey),
	).With(log.String(myTraceKey, "123674523547263541672341874"))
	defer logger.Sync()

	logger.Info("hello", log.String("test_trace", "nothing"))
}
