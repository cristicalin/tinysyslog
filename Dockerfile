FROM golang:alpine
MAINTAINER Alexandre Ferland <aferlandqc@gmail.com>

RUN apk add --no-cache git && mkdir -p /data/logs

WORKDIR /go/src/github.com/clearbit/tinysyslog

ADD . /go/src/github.com/clearbit/tinysyslog

RUN go-wrapper install

EXPOSE 5140 5140/udp

CMD ["/go/bin/tinysyslog", "--address", "0.0.0.0:5140", "--filesystem-filename", "/data/logs/syslog.log", "--log-file", "stdout"]
