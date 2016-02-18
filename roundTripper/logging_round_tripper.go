package roundTripper

import (
	"errors"
	"log"
	"net/http"
	"time"
)

type LoggingRoundTripper struct {
	transport http.RoundTripper
}

func NewLoggingRoundTripper() *LoggingRoundTripper {
	return &LoggingRoundTripper{
		transport: http.DefaultTransport,
	}
}

// forward to the URL
// Send response back to the Router

func (lrt *LoggingRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	var err error
	var res *http.Response
	start := time.Now()
	if request.Host == "No Host" {
		return nil, errors.New("Incoming Request Invalid")
	}
	res, err = lrt.transport.RoundTrip(request)
	if err != nil {
		return nil, err
	}
	res.Header.Add("X-Response-Forwarding", res.Status)
	log.Printf("Time Elapsed RoundTrip %v", time.Since(start))
	return res, err
}
