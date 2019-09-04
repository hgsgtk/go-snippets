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

`os` パッケージの`os.Stdin`から標準入力を受け取ることができます。
`bufio`パッケージの`bufio.Scanner`を使用することで一行ずつ読み取っていくことができます。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		fmt.Println(stdin.Text() + " by Go")
	}
}
```

これをパイプつなぎで実行してみましょう。

```bash
$ echo huga | go run stdin/main.go
huga by Go
```

標準入力で受け取った huga が ` by Go`という文字列と結合されて出力される

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

## ファイルを扱う
ファイルを読み書きしてみましょう。まずは、 `hoge.txt` ファイルがあるとして、それを開いてみます。

### ファイルを読む
```bash
// hoge.txtを用意
$ echo hoge > hoge.txt
```

このファイルの中身を読んで出力してみましょう。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("hoge.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ft := bufio.NewScanner(f)
	for ft.Scan() {
		fmt.Println(ft.Text())
	}
}
```

ファイルを開くには `os.Open()` という関数を使用できます。
また、`defer`という式が現れました。これは遅延実行のための式で、`defer f.Close()`はこの関数終了時に実行されます。
つまり、main関数を一通り実行し終わったタイミングで実行されます。
Goでは、このように「ファイルを開いたよ」に対する後片付け的な処理を直後に`defer`で宣言しておくことで表現することが多いです。

ここで、またしても `bufio.NewScanner()`が現れました。先程は標準入力を受け取って、今回はファイルを受け取っています。
なぜ、このように汎用的なことができているかということをちょっとだけ深堀りします。

### 入出力の抽象化
Goでは`io.Reader`・`io.Writer`という２つのInterface型を提供しています。

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}

type Reader interface {
        Read(p []byte) (n int, err error)
}
```

https://golang.org/pkg/io/#Reader
https://golang.org/pkg/io/#Writer

たった一つのメソッドをもつシンプルなインターフェース
このインターフェースによって、ファイル・標準入出力・ネットワーク・メモリなどが透過的に扱えるように設計されています。

今回の例であれば、`bufio.NewScanner`ですがこの実装の中身はこうなっています。

```go
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:            r,
		split:        ScanLines,
		maxTokenSize: MaxScanTokenSize,
	}
}
```

`io.Reader`を引数に求めているので、さまざまな入力にたいして扱えるものになっているというわけです。

### ファイルを書き込む
次にファイルを書き込んでみましょう。

```go
package main

import "os"

func main() {
	f, err := os.Create("huga.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

```

ファイルを作成するには`os.Create`を利用します。これを実行すると、`huga.txt`が作成されているでしょう。
