FROM golang:1.25-alpine as gotools

RUN mkdir /build

COPY . /build/gocloud-blob

RUN apk update && apk upgrade 

RUN cd /build/gocloud-blob \
    && go build -mod vendor -ldflags="-s -w" -o /usr/local/bin/copy-uri cmd/copy-uri/main.go

FROM alpine

RUN apk update && apk upgrade \
    && apk add git ca-certificates

COPY --from=gotools /usr/local/bin/copy-uri /usr/local/bin/
