# auth0-app
- auth0 golang tutorial
    - https://auth0.com/docs/quickstart/webapp/golang
- use mux as a router
- use negroni as a middleware

## Getting Started
- install docker(e.g. Docker for Mac)
- create auth0 apps
- set auth0 id to .env
- build application

```$xslt
$ make build
```

- run application

```$xslt
$ make run
```

- access localhost
    - http://localhost:3000

## Ref
- https://golang.org/pkg/os/#Getwd
- https://github.com/js-cookie/js-cookie
- https://github.com/urfave/negroni