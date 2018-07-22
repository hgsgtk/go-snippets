# db-design
sql.DBをどう扱うかについての調査

## Setup
- create database and run seed.sql

## Exec
- start server

```
$ go run main.go

$ curl -i localhost:3000/books/show?isbn=978-1503261969
$ curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5900" localhost:3000/books/create

```

## Todo
- [x] basic implement: dbを扱うhttpサービスの実装


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