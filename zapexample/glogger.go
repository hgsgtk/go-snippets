package zapexample

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() {
	// NewAtomicLevel creates an AtomicLevel with InfoLevel
	atom := zap.NewAtomicLevel()

	// production used な encoding configurationを取得できる
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.LevelKey = "type" // エラーレベルのキーを type にしたりできる
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// zap.Newでは柔軟にLoggerを生成できる方法
	// configとかをいろいろ変えるにはいい方法
	// NewCore creates a Core that writes logs to a WriteSyncer.
	bl := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg), // JSONエンコーディングするエンコーダー
		zapcore.Lock(os.Stdout),
		atom,
	))

	l := bl.With(zap.String("out", "stdout"))

	zap.ReplaceGlobals(l)
}
