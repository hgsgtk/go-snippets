package zapexample_test

import (
	"testing"

	"github.com/hgsgtk/go-snippets/zapexample"
	"go.uber.org/zap"
)

func TestInit(t *testing.T) {
	// スクリプト実行時に初期化
	// zap package の global logger をcustomしたものに書き換える
	zapexample.Init()
	zap.L().Error("hoge")
}

func TestNormalGLogger(t *testing.T) {
	// zapデフォルトのloggerもglobalで使えるよ
	// It's safe for concurrent use.
	zap.L().Error("hoge")
}
