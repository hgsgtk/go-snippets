# CLI Hands on!
Go言語でCLIを作成するためのあれこれをやってみるHands onです。基本的な構成要素として次のことを学んでみます。

- プログラム引数を扱う
- オプションを扱う
- 標準入力を受け取る
- コマンドを実行する
- ファイルを扱う

## プログラム引数を扱う
プログラム引数とはプログラム実行時に渡される引数です。たとえば、

```bash
cat hoge.txt
```

というコマンドであれば hoge.txt がそれに該当します。プログラム引数受け取れたらなんとなく最低限のCLIがつくれそうですよね。やってみましょう。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
}
```

実際に実行してみましょう。

```bash
go run main.go hoge hoge hoge
```

## オプションを扱う
コマンドっぽい何かを作る際は `-go=hoge` みたいにオプション扱いたいですよね。Goでは標準パッケージ flag でそれのニーズに答えることができます。

https://golang.org/pkg/flag/

