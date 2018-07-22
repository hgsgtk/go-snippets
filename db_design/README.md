# db-design
sql.DBをどう扱うかについての調査

## Setup
- create database and run seed.sql

## Folders
- basic: リファクタリング前のワンライナーコード
- globalvariable: db.SQLをパッケージ変数として扱う
- di: 依存性注入
- usingif: interfaceの活用
- usingctx: ctxにsql.DBをつめつめ

### globalvariable
- db.SQLをパッケージ変数として扱う
- すべてのdatabaseを扱う処理は同じmodelsパッケージ内に集約する
- パッケージ変数として扱っているためテストの並列実行は不可。

### di
- 引数もしくはユーザー型内に含め引き回す
- すべてのhandlersが同じパッケージ内にある前提。
- すべてのhandlersに同じ`Env`structがいる
- とはいえ、構造体を引数に求めているので別の実装をmockに使うなどは難しい

#### Refs
- configパッケージを間に挟むパターン
    - https://gist.github.com/alexedwards/8b4b0cd4495d7c3abadd
    - configという形でdbのglobal connectionを使い回

### usingif
- interfaceを活用する
- interfaceを挟むことによって、dbにアクセスする箇所をmock実装に差し替えることができる
- `GetAllBooks`とかをメソッドに持つInterfaceにしてるけどちゃんと考えないと無限にInterfaceが大きくなって本末転倒になるのは注意

### usingctx
- contextにdbを詰め詰めする
- （なんか怖い）

#### Refs
- [Contextアンチパターン](https://speakerdeck.com/timakin/contextantihatan)

## Exec
- start server

```
$ go run main.go

$ curl -i localhost:3000/books/show?isbn=978-1503261969
$ curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5900" localhost:3000/books/create

```

## Tips
- `sql.Open`はconnectionをcheckしてくれるわけではないので、`db.Ping`で実際にconnectionが通っているかを確認するとよい。
- `Request.FormValue()` method to fetch the value from the request query string.
- parameterの値をキーにfetchする場合は、prepared statementを作ることによってescapeしてくれるのでSQL injection対策となる
- `row.Scan`はレコードがない場合は、`ErrNoRows`というerror型を返すのでエラーハンドリングする
- https://github.com/jmoiron/sqlx

## Ref
- [Practical Persistence in Go: SQL Databases](https://www.alexedwards.net/blog/practical-persistence-sql)
- [Practical Persistence in Go: Organising Database Access](https://www.alexedwards.net/blog/organising-database-access)
- [Goを仕事で書き始める前に読んでおくといいドキュメント](https://qiita.com/Khigashiguchi/items/086947e93f565e755996)