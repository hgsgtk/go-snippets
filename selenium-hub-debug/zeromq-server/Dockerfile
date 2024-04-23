FROM golang:1.22.1-alpine

WORKDIR /app

RUN apk update
RUN apk add --no-cache alpine-sdk build-base gcc git libzmq pkgconfig zeromq-dev

RUN pkg-config --modversion libzmq

ENV GO111MODULE=on
ENV CGO_ENABLED=1

COPY go.mod .
COPY go.sum .

COPY . .

RUN go mod download

RUN go build -o server .

CMD ["./server"]
