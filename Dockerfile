FROM golang:1.14.1-buster

WORKDIR /code

RUN sed -i "s/deb.debian.org/mirrors.aliyun.com/" /etc/apt/sources.list
RUN sed -i "s/security.debian.org/mirrors.aliyun.com/" /etc/apt/sources.list

RUN apt update && apt install -y librdkafka-dev

COPY . .

ENV GOPROXY=https://goproxy.io

RUN go build -o ./bin/filtercmpt ./cmd/filtercmpt/hello.go

CMD ["/code/bin/filtercmpt"]

