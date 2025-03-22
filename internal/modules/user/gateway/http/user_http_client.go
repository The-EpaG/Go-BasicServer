package http

import (
	easyhttpclient "github.com/The-EpaG/EasyHttpClient"
)

type UserHttpClient interface {
	easyhttpclient.HttpClient
}
