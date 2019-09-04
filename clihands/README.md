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

```go
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
	fmt.Println(*msg) // *でポインタの値をとってる
	fmt.Println(strings.Repeat("にゃーん", nyan))
}
```

実際にオプションをつけて実行してみましょう。

```bash
$ go run main.go -msg="HelloVanillaJS" -nyan=2
  HelloVanillaJS
  にゃーんにゃーん
```

ちなみに、 `-h` や `-help` と渡すと、HELPを出すことができます。

```bash
$ go run rcvflg/main.go -h
Usage of /var/folders/lg/rdr0tvnd6kzblb0y1xmpvvx00000gn/T/go-build919209028/b001/exe/main:
  -msg string
        表示したいメッセージを入力してね (default "Hello PHPer")
  -nyan int
        にゃーんと言ってほしい回数 (default 1)
exit status 2
```

## 標準入力を受け取る
次は標準入力を受け取るパターンをやってみましょう。標準入力を受け取るとパイプで受け取ってなにかすることができますね

## コマンドを実行する
`echo`など、osコマンドを実行できたらシェルスクリプトの代わりにもなりそうですね、チョット複雑であればGoのほうが良いかもしれませんね。

osコマンドを実行するには、 `os/exec` パッケージを使います。

https://golang.org/pkg/os/exec/

簡単なもので、 `echo Hello` をGoから実行してみましょう。

```go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "Hello")
	output, err := cmd.Output()
	if err != nil {
		panic(err) // 稼働中の本番サービスなどではpanicしない
	}
	fmt.Print(string(output))
}
```

