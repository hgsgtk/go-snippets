# sample of logger using zap

## patterns
- global logger

## 解説
### zapのWriteSyncer interface とは
io.Writer interfaceをもち、さらにSync() をシグネチャにもつもの。`os.File/os.Stderr/os.Stdout`はこの辺をもっている

```go
// A WriteSyncer is an io.Writer that can also flush any buffered data. Note
// that *os.File (and thus, os.Stderr and os.Stdout) implement WriteSyncer.
type WriteSyncer interface {
	io.Writer
	Sync() error
}
```
### global logger
encodingのオプションをいろいろいじれます。以下、zapのproduction encoding level

```go
	func NewProductionEncoderConfig() zapcore.EncoderConfig {
		return zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
	}
```

デフォルトでくるキーもいろいろ変えれるよ

