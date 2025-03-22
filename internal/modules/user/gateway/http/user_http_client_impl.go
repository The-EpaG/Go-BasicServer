package http

import (
	"github.com/The-EpaG/EasyHttpClient"
	"github.com/The-EpaG/EasyHttpClient/transport"
	"github.com/The-EpaG/Go-BasicServer/internal/middleware/interceptor"
	"github.com/The-EpaG/Go-BasicServer/pkg/config"
)

type UserHttpClientImpl struct {
	*easyhttpclient.HttpClientImpl
}

func NewUserHttpClientImpl() UserHttpClient {
	cfg := config.Get()
	return &UserHttpClientImpl{
		HttpClientImpl: easyhttpclient.NewHttpClient(
			&easyhttpclient.HttpClientOptions{
				RequestInterceptors: []transport.EasyHttpClientRequestInterceptor{
					interceptor.NewLoggingRequestInterceptor(),
				},
				ResponseInterceptors: []transport.EasyHttpClientResponseInterceptor{
					interceptor.NewLoggingResponseInterceptor(),
				},
				BaseUrl: cfg.ExtService.BaseUrl,
			}),
	}
}
