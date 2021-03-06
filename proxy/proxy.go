package proxy

import (
	"bytes"
	"github.com/Pivotal-Japan/service-test/headers"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func NewReverseProxy(transport http.RoundTripper) *httputil.ReverseProxy {

	reverseProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			start := time.Now()
			RouterServiceheader := headers.NewRouteServiceHeaders()

			err := RouterServiceheader.ParseHeadersAndClean(&req.Header)

			if RouterServiceheader.IsValidRequest() && err == nil {
				body, err := ioutil.ReadAll(req.Body)
				if err != nil {
					log.Fatalln(err.Error())
				}
				req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				req.URL = RouterServiceheader.ParsedUrl
				req.Host = RouterServiceheader.ParsedUrl.Host
			} else {
				req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte{}))
				req.Host = "No Host"
				log.Print(RouterServiceheader.String())
			}

			log.Printf("Time Elapsed header %v ", time.Since(start))

		},
		Transport: transport,
	}
	return reverseProxy
}
