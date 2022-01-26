# HTTPS server in Go

https://medium.com/rungo/secure-https-servers-in-go-a783008b36da

## Generate a private key and a self-signed SSL certificate

ref: https://medium.com/geekculture/creating-a-local-websocket-server-with-tls-ssl-is-easy-as-pie-de1a2ef058e0

- Generate a `local_hgsgtk_com.ley` (private key) 

```bash
openssl genrsa -des3 -out local_hgsgtk_com.pem 2048
```

- Generate a `local_hgsgtk_com.csr` which is a certificate signing request (CSR).

```
openssl req -new -key local_hgsgtk_com.pem -out local_hgsgtk_com.csr
```

- Generate the `local_hgsgtk_com.crt` which is the self-signed certificate signed by own `localhost.key` private key.

```bash
openssl x509 -req -days 365 -in local_hgsgtk_com.csr -signkey local_hgsgtk_com.pem -out local_hgsgtk_com.crt
```

- Decrypt the RSA private key

```
openssl rsa -in local_hgsgtk_com.pem -out local_hgsgtk_com.key
```
