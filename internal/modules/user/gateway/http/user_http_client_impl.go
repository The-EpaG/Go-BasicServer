package http

import (
	"github.com/The-EpaG/Go-BasicServer/internal/middleware"
	"github.com/The-EpaG/Go-BasicServer/internal/middleware/interceptor"
	"github.com/The-EpaG/Go-BasicServer/pkg/config"
	defaultHttp "github.com/The-EpaG/Go-BasicServer/pkg/http"
)

type UserHttpClientImpl struct {
	*defaultHttp.HttpClientImpl
}

func NewUserHttpClientImpl() UserHttpClient {
	cfg := config.Get()
	return &UserHttpClientImpl{
		HttpClientImpl: defaultHttp.NewHttpClientImpl(
			&defaultHttp.HttpClientOptions{
				RequestInterceptors: []middleware.HttpClientRequestInterceptor{
					interceptor.LoggingRequestInterceptor,
				},
				ResponseInterceptors: []middleware.HttpClientResponseInterceptor{
					interceptor.LoggingResponseInterceptor,
				},
				BaseUrl: cfg.ExtService.BaseUrl,
			}),
	}
}
