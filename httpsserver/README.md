# HTTPS server in Go

https://medium.com/rungo/secure-https-servers-in-go-a783008b36da

## Generate a private key and a self-signed SSL certificate

- Generate a `localhost.ley` (private key) and `localhost.csr`.

```bash
openssl req -new -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr
```

- Generate the `localhost.crt` which is the self-signed certificate signed by own `localhost.key` private key.

```bash
openssl x509 -req -days 365 -in localhost.csr -signkey localhost.key -out localhost.crt
```

