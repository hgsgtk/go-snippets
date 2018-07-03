# layerd-basic
レイヤードアーキテクチャになんとなく沿う

## パッケージ構成
- src: application code
    - presenter
        - api
            - server
                - auth
                - handler
                - middleware
                - router?
    - usecase
    - domain: business logic
        - model: store model object struct
        - service: domain service
        - repository: database handler (especialy CRUD)
    - infrastructure
        - persistence
            - datastore
        - network
            - mail
            - aws
            - flentd
            - api
            ...etc?
    - registry: DIP
- config: config
- vendor: libraries

## ライブラリ
- routing
    - gorilla/mux
- middleware
    - urfave/negroni

## レイヤードアーキテクチャ
- https://speakerdeck.com/sonatard/next-currency-gaego?slide=25

### そもそもレイヤードアーキテクチャとは

### 想定している利点
- business logicに対するテストを重点的にやりたい場合、レイヤーは疎結合にしておく１つの指針になりうるか
- cakephpのfat controllerになりたくない

## application
- https://speakerdeck.com/sonatard/next-currency-gaego?slide=29
    - domain model, infra(db), 外部APIなどの取りまとめ
    - ここにはビジネスロジックは書かない
    - domain modelの利用者

### 参考
- Goのパッケージ構成の失敗遍歴と現状確認 by @timakinさん
    - https://medium.com/@timakin/go%E3%81%AE%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E6%A7%8B%E6%88%90%E3%81%AE%E5%A4%B1%E6%95%97%E9%81%8D%E6%AD%B4%E3%81%A8%E7%8F%BE%E7%8A%B6%E7%A2%BA%E8%AA%8D-fc6a4369337
- pospomeさんレイヤについて
    - https://www.slideshare.net/pospome/go-80591000
- recruit lifestyle package構成について
    - https://engineer.recruit-lifestyle.co.jp/techblog/2018-03-16-go-ddd/
- ドメインモデル貧血症
    - http://bliki-ja.github.io/AnemicDomainModel/
- go clent architecute sample project
    - https://hackernoon.com/golang-clean-archithecture-efd6d7c43047
    - https://github.com/bxcodec/go-clean-arch
- https://qiita.com/little_hand_s/items/ebb4284afeea0e8cc752
- https://qiita.com/oshiro/items/65d108e533a36c87a6da
- clean architecture sample
    - https://github.com/ktr0731/cris/tree/master/server
- https://golang.org/pkg/context/
    - Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx:
- middlewareについて
    - https://mattstauffer.com/blog/laravel-5.0-middleware-filter-style/

### その他

- mux & negroni

```go
// Run run http server
func Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/users/", handler.IndexHandler)

	n := negroni.Classic()
	n.UseHandler(r)
	return http.ListenAndServe(":8080", n)
}
```

- httpauth: basic認証

```go
// Run run http server
func Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/users/", handler.IndexHandler)

	http.Handle("/", httpauth.SimpleBasicAuth("test", "test")(r))
	return http.ListenAndServe(":8080", nil)
}
```

-> % curl -u test:test http://localhost:8080/users/
Hello, http server.