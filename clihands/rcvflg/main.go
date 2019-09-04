package main

import (
	"flag"
	"fmt"
	"strings"
)

var msg = flag.String("msg", "Hello PHPer", "表示したいメッセージを入力してね")
var nyan int

// パッケージの初期化関数です
func init() {
	// & はポインタ
	flag.IntVar(&nyan, "nyan", 1, "にゃーんと言ってほしい回数")
}

func main() {
	flag.Parse() // 実際にフラグを設定します
	fmt.Println(msg)
	fmt.Println(strings.Repeat("にゃーん", nyan))
}
