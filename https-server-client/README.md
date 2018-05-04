# https-client-server
HTTPS client app and server app.

## Getting Started
### Prepare certification
- create secret key

```bash
openssl genrsa -out ca.key 2048
```

- create csr

```bash
openssl req -new -sha256 -ket ca.key -out ca.csr -config openssl.cnf
```

- create certification file

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

- create server certification file

```bash
openssl x509 -req -days 365 -in server.csr -sha256 -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile {openssl.cnf path} -extensions Server
```

### Prepare client certification
- create client secret key

```bash
openssl genrsa -out client.key 2048
```

- create client csr file

```bash
openssl req -new -nodes -sha256 -key client.key -out client.csr -config {openssl.cnf path}
```

- create client certification

```bash
openssl x509 -req -days 365 -in client.csr -sha256 -out client.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile {openssl.cnf file path} -extensions Client
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
## Version
- basic
    - https://github.com/Khigashiguchi/go_snippet/pull/2/commits/c6e862abba408ac5558f4e2e172a06f2e47632dd
- server require client certification

## Ref
- https://www.oreilly.co.jp/books/9784873118048/
