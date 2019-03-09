Go言語でアプリケーションを実装し、それに対してユニットテストを書く場合、**テーブル駆動テスト**の技法を活用するケースは多いです。  
今回は、そのテーブル駆動テストのテーブルをちょっと見通しよくするためのちょっとした工夫をご紹介。  
見やすさは非常に主観的と思うので採用可否はお任せ。

例えば、次のようなFizzBuzzのコードがあるとします。

```go
package fizzbuzz

import "strconv"

// 数値に合わせてFizzBuzz/Fizz/Buzzを返す
func Run(num int) string {
	var res string
	switch {
	case num%15 == 0:
		res = "FizzBuzz"
	case num%5 == 0:
		res = "Buzz"
	case num%3 == 0:
		res = "Fizz"
	default:
		res = strconv.Itoa(num)
	}
	return res
}
```

# テーブル駆動テストをやっていく

テーブル駆動テストを書いていきます。よく見られる書き方として、テストテーブルを構造体の配列（`[]struct）`）として定義してサブテストなどで実行するケースです。

```go
func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected string
	}{
		{
			name:     "15で割り切れる場合FizzBuzz",
			num:      45,
			expected: "FizzBuzz",
		},
		{
			name:     "5で割り切れる場合Buzz",
			num:      40,
			expected: "Buzz",
		},
		{
			name:     "3で割り切れる場合Buzz",
			num:      39,
			expected: "Fizz",
		},
		{
			name:     "15,5,3で割り切れない場合そのまま",
			num:      37,
			expected: "37",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := fizzbuzz.Run(tt.num); tt.expected != actual {
				t.Errorf("Run is expected '%s', but got '%s'", tt.expected, actual)
			}
		})
	}
}
```

このテストを実行すると次のような結果になりますね。

```shell
=== RUN   TestRun
--- PASS: TestRun (0.00s)
=== RUN   TestRun/15で割り切れる場合FizzBuzz
    --- PASS: TestRun/15で割り切れる場合FizzBuzz (0.00s)
=== RUN   TestRun/5で割り切れる場合Buzz
    --- PASS: TestRun/5で割り切れる場合Buzz (0.00s)
=== RUN   TestRun/3で割り切れる場合Buzz
    --- PASS: TestRun/3で割り切れる場合Buzz (0.00s)
=== RUN   TestRun/15,5,3で割り切れない場合そのまま
    --- PASS: TestRun/15,5,3で割り切れない場合そのまま (0.00s)
PASS
```

# ちょっと見やすくする
個人的な主観ですが、テーブルが大きくなったりケース数が増えてくると、次のような課題感を持ちました。

- 「どのようなケースがあるのか」を見る際に、テーブルのケース名と中身が同じ構造体内に同列に書かれるので、一覧しにくい。

なので、`map[string]struct{}`の形でケース名と中身を分離するちょっとした工夫をしてみています。この形式にすると次のようなテストコードになります。

```go
package fizzbuzz_test

import (
	"github.com/hgsgtk/go-snippets/fizzbuzz"
	"testing"
)

func TestRun(t *testing.T) {
	tests := map[string]struct {
		num      int
		expected string
	}{
		"15で割り切れる場合FizzBuzz": {
			num:      45,
			expected: "FizzBuzz",
		},
		"5で割り切れる場合Buzz": {
			num:      40,
			expected: "Buzz",
		},
		"3で割り切れる場合Buzz": {
			num:      39,
			expected: "Fizz",
		},
		"15,5,3で割り切れない場合そのまま": {
			num:      37,
			expected: "37",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if actual := fizzbuzz.Run(tt.num); tt.expected != actual {
				t.Errorf("Run is expected '%s', but got '%s'", tt.expected, actual)
			}
		})
	}
}
```

ケース名がわかりやすくなったかと思います。もし、同じような課題感を持っていた方がいらっしゃれば試してみてください。

以上