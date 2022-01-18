# Proxy which handles CONNECT request (HTTPS)

```bash
$ go run main.go
2022/01/18 19:18:18 Starting proxy server on 127.0.0.1:8080


2022/01/18 19:18:21 client: 127.0.0.1:52989, method: CONNECT, url: //google.com:443, host: google.com:443
2022/01/18 19:18:21 headers: http.Header{"Proxy-Connection":[]string{"Keep-Alive"}, "User-Agent":[]string{"curl/7.77.0"}}
2022/01/18 19:18:21 target site remote address: 216.58.220.110:443
2022/01/18 19:18:21 client: 127.0.0.1:52991, method: CONNECT, url: //www.google.com:443, host: www.google.com:443
2022/01/18 19:18:21 headers: http.Header{"Proxy-Connection":[]string{"Keep-Alive"}, "User-Agent":[]string{"curl/7.77.0"}}
2022/01/18 19:18:22 target site remote address: 142.251.42.196:443
```

```bash
$ curl -I -Lv -x http://127.0.0.1:8080 \
https://google.com
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
* allocate connect buffer!
* Establish HTTP proxy tunnel to google.com:443
> CONNECT google.com:443 HTTP/1.1
> Host: google.com:443
> User-Agent: curl/7.77.0
> Proxy-Connection: Keep-Alive
>
< HTTP/1.0 200 OK
HTTP/1.0 200 OK
<

* Proxy replied 200 to CONNECT request
* CONNECT phase completed!
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*  CAfile: /etc/ssl/cert.pem
*  CApath: none
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-ECDSA-CHACHA20-POLY1305
* ALPN, server accepted to use h2
* Server certificate:
*  subject: CN=*.google.com
*  start date: Dec  8 21:28:49 2021 GMT
*  expire date: Mar  2 21:28:48 2022 GMT
*  subjectAltName: host "google.com" matched cert's "google.com"
*  issuer: C=US; O=Google Trust Services LLC; CN=GTS CA 1C3
*  SSL certificate verify ok.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x14a810600)
> HEAD / HTTP/2
> Host: google.com
> user-agent: curl/7.77.0
> accept: */*
>
< HTTP/2 301
HTTP/2 301
< location: https://www.google.com/
location: https://www.google.com/
< content-type: text/html; charset=UTF-8
content-type: text/html; charset=UTF-8
< date: Tue, 18 Jan 2022 10:18:21 GMT
date: Tue, 18 Jan 2022 10:18:21 GMT
< expires: Thu, 17 Feb 2022 10:18:21 GMT
expires: Thu, 17 Feb 2022 10:18:21 GMT
< cache-control: public, max-age=2592000
cache-control: public, max-age=2592000
< server: gws
server: gws
< content-length: 220
content-length: 220
< x-xss-protection: 0
x-xss-protection: 0
< x-frame-options: SAMEORIGIN
x-frame-options: SAMEORIGIN
< alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000,h3-Q050=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443"; ma=2592000,quic=":443"; ma=2592000; v="46,43"
alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000,h3-Q050=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443"; ma=2592000,quic=":443"; ma=2592000; v="46,43"

<
* Connection #0 to host 127.0.0.1 left intact
* Issue another request to this URL: 'https://www.google.com/'
* Hostname 127.0.0.1 was found in DNS cache
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#1)
* allocate connect buffer!
* Establish HTTP proxy tunnel to www.google.com:443
> CONNECT www.google.com:443 HTTP/1.1
> Host: www.google.com:443
> User-Agent: curl/7.77.0
> Proxy-Connection: Keep-Alive
>
< HTTP/1.0 200 OK
HTTP/1.0 200 OK
<

* Proxy replied 200 to CONNECT request
* CONNECT phase completed!
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*  CAfile: /etc/ssl/cert.pem
*  CApath: none
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS change cipher, Change cipher spec (1):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-ECDSA-CHACHA20-POLY1305
* ALPN, server accepted to use h2
* Server certificate:
*  subject: CN=www.google.com
*  start date: Dec  8 22:50:34 2021 GMT
*  expire date: Mar  2 22:50:33 2022 GMT
*  subjectAltName: host "www.google.com" matched cert's "www.google.com"
*  issuer: C=US; O=Google Trust Services LLC; CN=GTS CA 1C3
*  SSL certificate verify ok.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x14a810600)
> HEAD / HTTP/2
> Host: www.google.com
> user-agent: curl/7.77.0
> accept: */*
>
< HTTP/2 200
HTTP/2 200
< content-type: text/html; charset=ISO-8859-1
content-type: text/html; charset=ISO-8859-1
< p3p: CP="This is not a P3P policy! See g.co/p3phelp for more info."
p3p: CP="This is not a P3P policy! See g.co/p3phelp for more info."
< date: Tue, 18 Jan 2022 10:18:22 GMT
date: Tue, 18 Jan 2022 10:18:22 GMT
< server: gws
server: gws
< x-xss-protection: 0
x-xss-protection: 0
< x-frame-options: SAMEORIGIN
x-frame-options: SAMEORIGIN
< expires: Tue, 18 Jan 2022 10:18:22 GMT
expires: Tue, 18 Jan 2022 10:18:22 GMT
< cache-control: private
cache-control: private
< set-cookie: 1P_JAR=2022-01-18-10; expires=Thu, 17-Feb-2022 10:18:22 GMT; path=/; domain=.google.com; Secure
set-cookie: 1P_JAR=2022-01-18-10; expires=Thu, 17-Feb-2022 10:18:22 GMT; path=/; domain=.google.com; Secure
< set-cookie: NID=511=SOoAK3H0Krlj8ckp4ny19INeRRUWH-IzYLGtQARjHQONfegazPBR3qnb6MlaMWW7U3pY_hj2k68N_7FP0FeK0JY39YY37zJIz3cV_Q0QznaXWsLv-r2Ft17ucHA1mhXDwDfE2PBbJqEmfSyW6NVWdo9woZCGmpJup-BqzVqCsSc; expires=Wed, 20-Jul-2022 10:18:22 GMT; path=/; domain=.google.com; HttpOnly
set-cookie: NID=511=SOoAK3H0Krlj8ckp4ny19INeRRUWH-IzYLGtQARjHQONfegazPBR3qnb6MlaMWW7U3pY_hj2k68N_7FP0FeK0JY39YY37zJIz3cV_Q0QznaXWsLv-r2Ft17ucHA1mhXDwDfE2PBbJqEmfSyW6NVWdo9woZCGmpJup-BqzVqCsSc; expires=Wed, 20-Jul-2022 10:18:22 GMT; path=/; domain=.google.com; HttpOnly
< alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000,h3-Q050=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443"; ma=2592000,quic=":443"; ma=2592000; v="46,43"
alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000,h3-Q050=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443"; ma=2592000,quic=":443"; ma=2592000; v="46,43"
```
