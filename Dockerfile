FROM golang:latest
WORKDIR /app
RUN go get gopkg.in/DataDog/dd-trace-go.v1/ddtrace
RUN go get github.com/DataDog/datadog-go/statsd
RUN go get "golang.org/x/xerrors"
RUN go get "golang.org/x/time/rate"
COPY . .
RUN go build -o main .
WORKDIR /dist
RUN cp /app/main .
CMD ["/dist/main"]