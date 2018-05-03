# github-oauth-cli
## github rest api v3

- schema

```
-> % curl -i https://api.github.com/users/octocat/orgs
HTTP/1.1 200 OK
Date: Thu, 03 May 2018 13:04:08 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 5
Server: GitHub.com
Status: 200 OK
X-RateLimit-Limit: 60
X-RateLimit-Remaining: 59
X-RateLimit-Reset: 1525356248
Cache-Control: public, max-age=60, s-maxage=60
Vary: Accept
ETag: "98f0c1b396a4e5d54f4d5fe561d54b44"
X-GitHub-Media-Type: github.v3; format=json
Access-Control-Expose-Headers: ETag, Link, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval
Access-Control-Allow-Origin: *
Strict-Transport-Security: max-age=31536000; includeSubdomains; preload
X-Frame-Options: deny
X-Content-Type-Options: nosniff
X-XSS-Protection: 1; mode=block
Referrer-Policy: origin-when-cross-origin, strict-origin-when-cross-origin
Content-Security-Policy: default-src 'none'
X-Runtime-rack: 0.020166
Vary: Accept-Encoding
X-GitHub-Request-Id: C85C:20CB:B99F0C:ECB04A:5AEB08C7

[

]
```

## Reference
- https://www.oreilly.co.jp/books/9784873118048/
    - 11章：クライアント視点で見るRestful API
