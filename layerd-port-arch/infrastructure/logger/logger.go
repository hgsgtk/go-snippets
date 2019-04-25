package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Writer specifies output of logger.
var Writer zapcore.WriteSyncer = os.Stdout

// zapのlogger instanceを作る
func newLogger(writer zapcore.WriteSyncer) *zap.Logger {
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	bl := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(writer),
		atom,
	))

	// AWSのCloudWatchLogsでは、標準出力・標準エラー出力が同一ストリーム内に混ざって出力されるため、
	// 吐き出すログレベルでどちらに対して出力したかを明示しておく意図。
	l := bl.With(zap.String("out", "stdout"))
	return l
}

// Init replace global zap logger to custom logger.
// 環境ごとの設定差異があるのであればここでlogger生成時に指定。
func Init(output zapcore.WriteSyncer) {
	logger := newLogger(output)
	zap.ReplaceGlobals(logger)
}

// Logger return logger instance.
func Logger() *zap.Logger {
	return zap.L()
}
