FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
RUN go get gopkg.in/DataDog/dd-trace-go.v1/ddtrace
RUN go get github.com/DataDog/datadog-go/statsd
RUN go get "golang.org/x/xerrors"
CMD ["/app/main"]