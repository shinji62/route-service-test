package main

import (
	"github.com/Pivotal-Japan/service-test/proxy"
	"github.com/Pivotal-Japan/service-test/roundTripper"
	"log"
	"net/http"
)

func main() {
	roundTripper := roundTripper.NewLoggingRoundTripper()
	proxy := proxy.NewReverseProxy(roundTripper)
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
