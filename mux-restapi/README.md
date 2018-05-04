# mux-restapi
RestAPI writtern in go, using mux as a router.

## Endpoints
### GET /books

```
$ curl http://localhost:8000/books

[{"id":"1","isbn":"1234567","title":"Book One","author":{"firstname":"Taro","lastname":"Suzuki"}},{"id":"2","isbn":"8901234","title":"Book Two","author":{"firstname":"Hanako","lastname":"Yamada"}},{"id":"298498081","isbn":"123456","title":"Post Book","author":{"firstname":"Hiroshi","lastname":"Sato"}},{"id":"427131847","isbn":"123456","title":"Post Book","author":{"firstname":"Hiroshi","lastname":"Sato"}}]
```
### GET /books/:id

```
$ curl http://localhost:8000/books/3

{"id":"","isbn":"","title":"","author":null}
```
### POST /books

```
$ curl -H "Accept: application/json" -H "Content-Type: application/json" -d '{"isbn": "123456", "title": "Post Book", "author": {"firstname": "Hiroshi", "lastname": "Sato"}}' http://localhost:8000/books

{"id":"427131847","isbn":"123456","title":"Post Book","author":{"firstname":"Hiroshi","lastname":"Sato"}}
```

### PUT /books/:id

```
$ curl -X PUT -H "Accept: application/json" -H "Content-Type: application/json" -d '{"isbn": "123456", "title": "Modified Title", "author": {"firstname": "Hiroshi", "lastname": "Sato"}}' http://localhost:8000/books/298498081

{"id":"298498081","isbn":"123456","title":"Modified Title","author":{"firstname":"Hiroshi","lastname":"Sato"}}
```
### DELETE /books/:id

```
curl -X DELETE http://localhost:8000/books/1
```

## Ref
- [https://github.com/bradtraversy/go_restapi](https://github.com/bradtraversy/go_restapi)
- [https://github.com/gorilla/mux](https://github.com/gorilla/mux)
