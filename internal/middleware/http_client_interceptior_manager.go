package middleware

import (
	"log"
	"net/http"
)

type HttpClientRequestInterceptor = func(req *http.Request) (*http.Request, error)
type HttpClientResponseInterceptor = func(req *http.Response) (*http.Response, error)

type HttpClientInterceptorManagerRoundTripper struct {
	Proxied              http.RoundTripper
	RequestInterceptors  []HttpClientRequestInterceptor
	ResponseInterceptors []HttpClientResponseInterceptor
}

func (lrt *HttpClientInterceptorManagerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var err error
	for _, interceptor := range lrt.RequestInterceptors {
		req, err = interceptor(req)
		if err != nil {
			return nil, err
		}
	}

	resp, err := lrt.Proxied.RoundTrip(req)

	if err != nil {
		log.Printf("Error in roundtrip: %v\n", err)
		return nil, err
	}

	for _, interceptor := range lrt.ResponseInterceptors {
		resp, err = interceptor(resp)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
