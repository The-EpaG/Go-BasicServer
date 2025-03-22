package interceptor

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/The-EpaG/Go-BasicServer/pkg/logging"
)

type LoggingRequestInterceptor struct {}

func (i *LoggingRequestInterceptor) Intercept(req *http.Request) (*http.Request, error) {
	var requestBody []byte
	if req.Body != nil {
		requestBody, _ = io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(requestBody)) //Set the body to the new buffer
	}
	logging.Info(req.Context(), "Client Request: "+req.Method+" "+req.URL.String())
	logging.Info(req.Context(), "Client Request Payload: "+string(requestBody))

	return req, nil
}

func NewLoggingRequestInterceptor() *LoggingRequestInterceptor {
	return &LoggingRequestInterceptor{}
}

type LoggingResponseInterceptor struct {}

func (i *LoggingResponseInterceptor) Intercept(resp *http.Response) (*http.Response, error) {
	var responseBody []byte
	if resp.Body != nil {
		responseBody, _ = io.ReadAll(resp.Body)
		resp.Body = io.NopCloser(bytes.NewBuffer(responseBody))
	}
	logging.Info(context.Background(), "Client Received Response: "+resp.Status)
	logging.Info(context.Background(), "Client Received Response Payload: "+string(responseBody))

	return resp, nil
}

func NewLoggingResponseInterceptor() *LoggingResponseInterceptor {
	return &LoggingResponseInterceptor{}
}