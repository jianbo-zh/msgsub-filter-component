FROM golang:1.14.1-alpine3.11

WORKDIR /code

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories

RUN apk update && apk add librdkafka-dev pkgconf

COPY . .

ENV GOPROXY=https://goproxy.io

RUN go build -o ./bin/filtercmpt ./cmd/filtercmpt/hello.go

CMD ["/code/bin/filtercmpt"]

