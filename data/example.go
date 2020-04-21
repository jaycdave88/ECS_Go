package main

import (
	"log"
	"net/http"
	"strings"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// From the docs https://docs.datadoghq.com/tracing/setup/go/#configuration

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	// curl IP for ECS to set DD_TRACE_AGENT_PORT
	resp, err :=http.Get("http://169.254.169.254/latest/meta-data/local-ipv4")
    if err == nil {
        addr := net.JoinHostPort(
        resp,
        8126,
        )
        os.Setenv("DD_AGENT_HOST", addr)
        tracer.Start(tracer.WithAgentAddr(addr))
		tracer.Start(tracer.WithDebugMode(true))

        defer tracer.Stop()

	mux := httptrace.NewServeMux() // init the http tracer
	mux.HandleFunc("/", sayHello)  // use the tracer to handle the urls

	err := http.ListenAndServe(":9090", mux) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}