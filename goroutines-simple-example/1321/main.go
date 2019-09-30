package main

import (
	"fmt"
	"time"
)

// 第13章 Go言語と並列処理
// 並行 Concurrent: CPU数、コア数の限界を超えて複数の仕事を同時に行う
// 並列 Parallel: 複数CPU、コアを効率よく扱って計算速度を上げる

func sub1(c int) {
	fmt.Println("share by arguments:", c*c)
}

// 出力順番が実行ごとに変わる、実行順序は不定
func main() {
	// 引数で渡す
	go sub1(10)

	// クロージャのキャプチャ渡し
	// 無名関数に暗黙な引数が追加され、暗黙の引数にデータが渡される
	c := 20
	go func() {
		fmt.Println("share by capture", c*c)
	}()
	time.Sleep(time.Second)
}
