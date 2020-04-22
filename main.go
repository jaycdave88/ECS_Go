package main

import (
    "log"
    "net/http"
    "net"
    "os"
    "io/ioutil"
    "strings"

    httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
        message := r.URL.Path
        message = strings.TrimPrefix(message, "/")
        message = "Hello " + message
        w.Write([]byte(message))
}

func main() {
    	// curl IP for ECS to set DD_TRACE_AGENT_PORT
        port := "8126"
    	resp, err := http.Get("http://169.254.169.254/latest/meta-data/local-ipv4")
        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }
        host := string(bodyBytes)
    if err == nil {
        addr := net.JoinHostPort(host, port)
        os.Setenv("DD_AGENT_HOST", addr)
        tracer.Start(tracer.WithAgentAddr(addr))
        tracer.Start(tracer.WithDebugMode(true))

        defer tracer.Stop()

        mux := httptrace.NewServeMux()
        mux.HandleFunc("/", sayHello)

        err := http.ListenAndServe(":9090", mux)
        if err != nil{
            log.Fatal("ListenAndServe: ", err)
        }
    }    
}