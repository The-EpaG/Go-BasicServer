package http

import (
	"context"
)

type HttpClient interface {
	Get(ctx context.Context, path string) ([]byte, error)
	Post(ctx context.Context, path string, body []byte, contentType ...string) ([]byte, error)
	Put(ctx context.Context, path string, body []byte, contentType ...string) ([]byte, error)
	Delete(ctx context.Context, path string) ([]byte, error)
	Patch(ctx context.Context, path string, body []byte, contentType ...string) ([]byte, error)
}
