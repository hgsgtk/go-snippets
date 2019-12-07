# SAMでLambdaのローカル開発環境を作る

## Initialize

```go
$ sam init --runtime go1.x
Which template source would you like to use?
	1 - AWS Quick Start Templates
	2 - Custom Template Location
Choice: 1

Project name [sam-app]: sam-lamda-hello

Quick start templates may have been updated. Do you want to re-download the latest [Y/n]: Y

-----------------------
Generating application:
-----------------------
Name: sam-lamda-hello
Runtime: go1.x
Dependency Manager: mod
Application Template: hello-world
Output Directory: .

Next steps can be found in the README file at ./sam-lamda-hello/README.md
```

`sam init`で`--runtime`にて任意のランタイム言語を指定することでテンプレートを生成する。Go言語を利用する場合は、`go1.x`を利用すれば良い。

```
 .
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── hello-world                 <-- Source code for a lambda function
│   ├── main.go                 <-- Lambda function code
│   └── main_test.go            <-- Unit tests
└── template.yaml
```

テンプレートが作られると、丁寧なREADMEが現れるのでそれを見ていくと良い。が、それで終了というのも寂しいので、続きの手順も合わせて。

次に、`github.com/aws/aws-lamda-go`を`go get`する。

```
go get -u github.com/aws/aws-lambda-go/...
```

次に、GoコードをBuildする。`make build`が用意されてるのでそれを雑に実行すればよい。

```
$ make build
GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
```

## ローカルで起動する
　
```
$ sam local start-api
Mounting HelloWorldFunction at http://127.0.0.1:3000/hello [GET]
You can now browse to the above endpoints to invoke your functions. You do not need to restart/reload SAM CLI while working on your functions, changes will be reflected instantly/automatically. You only need to restart SAM CLI if you update your AWS SAM template
2019-12-07 22:13:23  * Running on http://127.0.0.1:3000/ (Press CTRL+C to quit)
```

こうすると、`http://127.0.0.1:3000/hello`からリクエストが受け付けられるようになる

試しに、curlで叩くとこうなります。

```
✗ curl http://127.0.0.1:3000/hello
Hello, xxx.xxx.xxx.xxx (IPアドレス)
```

その際に、`sam local start-api`の標準出力には次のように表示される

```
Fetching lambci/lambda:go1.x Docker container image......
Mounting /Users/hgsgtk/go/src/github.com/hgsgtk/go-snippets/sam-lamda-hello/hello-world as /var/task:ro,delegated inside runtime container
START RequestId: 47c68a41-5dcb-12dd-a2ea-87a3eb0f8cad Version: $LATEST
END RequestId: 47c68a41-5dcb-12dd-a2ea-87a3eb0f8cad
REPORT RequestId: 47c68a41-5dcb-12dd-a2ea-87a3eb0f8cad	Init Duration: 282.24 ms	Duration: 821.93 ms	Billed Duration: 900 ms	Memory Size: 128 MB	Max Memory Used: 26 MB
No Content-Type given. Defaulting to 'application/json'.
2019-12-07 22:16:28 127.0.0.1 - - [07/Dec/2019 22:16:28] "GET /hello HTTP/1.1" 200 -
```
