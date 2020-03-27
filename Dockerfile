FROM golang:1.14.1-buster

WORKDIR /code

RUN sed -i "s/deb.debian.org/mirrors.aliyun.com/" /etc/apt/sources.list
RUN sed -i "s/security.debian.org/mirrors.aliyun.com/" /etc/apt/sources.list

COPY . .

RUN cd ./dependence/librdkafka-1.3.0/ && \
    chmod u+x ./configure && \
    chmod u+x ./lds-gen.py && \
    ./configure --prefix /usr && \
    make && \
    make install && \
    cd /code

ENV GOPROXY=https://goproxy.io

RUN go build -o ./bin/filtercmpt ./cmd/filtercmpt/hello.go

CMD ["/code/bin/filtercmpt"]
