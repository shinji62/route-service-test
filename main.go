package main

import (
	"fmt"
	"github.com/Pivotal-Japan/service-test/headers"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
	"net/http"
)

// Forwards incoming requests to whatever location URL points to

func main() {

	fwd, _ := forward.New()

	redirect := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// let us forward this request to another server
		fmt.Printf("%s", req.Header)
		RouterServiceheader := headers.NewRouteServiceHeaders()
		RouterServiceheader.ParseHeaders(&req.Header)
		fmt.Printf("%s", RouterServiceheader.IsValidRequest())
		req.URL = testutils.ParseURI("http://localhost:63450")
		fwd.ServeHTTP(w, req)
	})

	// that's it! our reverse proxy is ready!
	s := &http.Server{
		Addr:    ":8080",
		Handler: redirect,
	}
	s.ListenAndServe()
}
