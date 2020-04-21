FROM golang:1.13.4

WORKDIR /home/

RUN apt-get update

RUN apt-get install -y nano

COPY data/example.go example.go

RUN go get gopkg.in/DataDog/dd-trace-go.v1/ddtrace

RUN go get github.com/DataDog/datadog-go/statsd

RUN go get "golang.org/x/xerrors"