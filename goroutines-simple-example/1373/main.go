package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("Initialize")
}

// 初期化処理を一度だけ行いたいときに使う
// init()という名前の関数が初期化関数として呼ばれるので、
// あえて必要になるまで遅延させたい場合に使える
var once sync.Once

func main() {
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
