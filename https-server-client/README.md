# https-client-server
HTTPS client app and server app.

## Getting Started
### Prepare client certification
- create client secret key

```bash
openssl genrsa -out ca.key 2048
```

- create client csr

```bash
openssl req -new -sha256 -ket ca.key -out ca.csr -config openssl.cnf
```

- create client certification file

```bash
openssl x509 -in ca.csr -days 365 -req -signkey ca.key -sha256 -out ca.crt -extfile {openssl.cnf path} -extensions CA
```

### Prepare server certification
- create server secret key

```bash
openssl genrsa -out server.key 2048
```

- create client csr

```bash

openssl req -new -nodes -sha256 -key server.key -out server.csr -config {openssl.cnf path}
```

Notice: answer common name 'localhost'

- create client certification file

```bash
openssl x509 -req -days 365 -in server.csr -sha256 -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile {openssl.cnf path} -extensions Server
```

## Run application
- Run server

```bash
go run server.go
```

- Execute Client

```bash
go run client.go
```

## Ref
- https://www.oreilly.co.jp/books/9784873118048/
