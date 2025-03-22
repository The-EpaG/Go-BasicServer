package mocks

import (
	"context"

	"github.com/The-EpaG/Go-BasicServer/pkg/http"
)

var (
	HttpGetReturnValue []byte
	HttpGetErrorValue  error = nil
)

type MockHttpClient struct {
	http.HttpClient
}

func (m *MockHttpClient) Get(ctx context.Context, url string) ([]byte, error) {
	return HttpGetReturnValue, HttpGetErrorValue
}
