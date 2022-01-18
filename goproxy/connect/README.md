# Connect proxy

## goproxy supports

- regular HTTP proxy
- HTTPS through CONNECT
- "hijacking" HTTPS connection using "Man in the Middle" style attack

## Initial version

```go
func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	log.Fatal(http.ListenAndServe(":8080", proxy))
}
```

Then, test it.

```bash
$ curl -I -Lv -x http://127.0.0.1:8080 \
https://example.com
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
* allocate connect buffer!
* Establish HTTP proxy tunnel to example.com:443
> CONNECT example.com:443 HTTP/1.1
> Host: example.com:443
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
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256
* ALPN, server accepted to use h2
* Server certificate:
*  subject: C=US; ST=California; L=Los Angeles; O=Verizon Digital Media Services, Inc.; CN=www.example.org
*  start date: Dec 10 00:00:00 2021 GMT
*  expire date: Dec  9 23:59:59 2022 GMT
*  subjectAltName: host "example.com" matched cert's "example.com"
*  issuer: C=US; O=DigiCert Inc; CN=DigiCert TLS RSA SHA256 2020 CA1
*  SSL certificate verify ok.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x13900b600)
> HEAD / HTTP/2
> Host: example.com
> user-agent: curl/7.77.0
> accept: */*
>
< HTTP/2 200
HTTP/2 200
< content-encoding: gzip
content-encoding: gzip
< accept-ranges: bytes
accept-ranges: bytes
< age: 504101
age: 504101
< cache-control: max-age=604800
cache-control: max-age=604800
< content-type: text/html; charset=UTF-8
content-type: text/html; charset=UTF-8
< date: Tue, 18 Jan 2022 06:11:18 GMT
date: Tue, 18 Jan 2022 06:11:18 GMT
< etag: "3147526947"
etag: "3147526947"
< expires: Tue, 25 Jan 2022 06:11:18 GMT
expires: Tue, 25 Jan 2022 06:11:18 GMT
< last-modified: Thu, 17 Oct 2019 07:18:26 GMT
last-modified: Thu, 17 Oct 2019 07:18:26 GMT
< server: ECS (sab/573E)
server: ECS (sab/573E)
< x-cache: HIT
x-cache: HIT
< content-length: 648
content-length: 648

<
* Connection #0 to host 127.0.0.1 left intact
```

- Log

```bash
$ go run main.go

2022/01/18 15:19:30 [004] INFO: Running 0 CONNECT handlers
2022/01/18 15:19:30 [004] INFO: Accepting CONNECT to example.com:443
```

### How it works

```go
func NewProxyHttpServer() *ProxyHttpServer {
	proxy := ProxyHttpServer{
		Logger:        log.New(os.Stderr, "", log.LstdFlags),
		reqHandlers:   []ReqHandler{},
		respHandlers:  []RespHandler{},
		httpsHandlers: []HttpsHandler{},
		NonproxyHandler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			http.Error(w, "This is a proxy server. Does not respond to non-proxy requests.", 500)
		}),
		Tr: &http.Transport{TLSClientConfig: tlsClientSkipVerify, Proxy: http.ProxyFromEnvironment},
	}
	proxy.ConnectDial = dialerFromEnv(&proxy)

	return &proxy
}
```

#### HttpsHandler

When a client send a CONNECT request to a host, 
the request is filtered through all the HttpsHandlers the proxy has.

```go
type HttpsHandler interface {
	HandleConnect(req string, ctx *ProxyCtx) (*ConnectAction, string)
}
```

https://pkg.go.dev/gopkg.in/elazarl/goproxy.v1#HttpsHandler

The HttpsHandler is used here.

```go
	todo, host := OkConnect, r.URL.Host
    // (omit)

	ctx.Logf("Running %d CONNECT handlers", len(proxy.httpsHandlers))
	todo, host := OkConnect, r.URL.Host
	for i, h := range proxy.httpsHandlers {
		newtodo, newhost := h.HandleConnect(host, ctx)

		// If found a result, break the loop immediately
		if newtodo != nil {
			todo, host = newtodo, newhost
			ctx.Logf("on %dth handler: %v %s", i, todo, host)
			break
		}
	}
```

https://github.com/elazarl/goproxy/blob/947c36da3153ff334e74d9d980de341d25f358ba/https.go#L80

#### Action

There are six kinds of connection:

1. ConnectAccept
2. ConnectReject
3. ConnectMitm
4. ConnectHijack
5. ConnectHTTPMitm
6. ConnectProxyAuthHijack

And, as actions using HTTPS narrows four: ConnectAccept, ConnectMitm, ConnectHTTPMitm, and ConnectReject.

```go
OkConnect       = &ConnectAction{Action: ConnectAccept, TLSConfig: TLSConfigFromCA(&GoproxyCa)}
MitmConnect     = &ConnectAction{Action: ConnectMitm, TLSConfig: TLSConfigFromCA(&GoproxyCa)}
HTTPMitmConnect = &ConnectAction{Action: ConnectHTTPMitm, TLSConfig: TLSConfigFromCA(&GoproxyCa)}
RejectConnect   = &ConnectAction{Action: ConnectReject, TLSConfig: TLSConfigFromCA(&GoproxyCa)}
```

https://pkg.go.dev/gopkg.in/elazarl/goproxy.v1#pkg-variables

```go
	switch todo.Action {
	case ConnectAccept:
		if !hasPort.MatchString(host) {
			host += ":80"
		}
		targetSiteCon, err := proxy.connectDial("tcp", host)
		if err != nil {
			httpError(proxyClient, ctx, err)
			return
		}
		ctx.Logf("Accepting CONNECT to %s", host)
```

https://github.com/elazarl/goproxy/blob/947c36da3153ff334e74d9d980de341d25f358ba/https.go#L102

By the way, goproxy contains a [X509Key pair](https://pkg.go.dev/crypto/tls#X509KeyPair).

- ConnectMitm: AlwaysMitm is a HttpsHandler that always eavesdrop https connections

```go
	/*
		$ curl -I -Lv -x http://127.0.0.1:8080 \
		https://example.com

		2022/01/18 15:47:00 [001] WARN: Cannot handshake client example.com:443 remote error: tls: unknown certificate authority
	*/
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
```

- AlwaysReject: a HttpsHandler that drops any CONNECT request

```go
	/*
		AlwaysReject is a HttpsHandler that drops any CONNECT request

		$ curl -I -Lv -x http://127.0.0.1:8080 \
		https://example.com

		2022/01/18 15:59:32 [001] INFO: Running 1 CONNECT handlers
		2022/01/18 15:59:32 [001] INFO: on 0th handler: &{1 <nil> 0x1022c7c60} example.com:443
	*/
	proxy.OnRequest().HandleConnect(goproxy.AlwaysReject)
```
