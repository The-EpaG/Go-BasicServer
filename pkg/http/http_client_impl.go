package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/The-EpaG/Go-BasicServer/internal/middleware"
	"github.com/The-EpaG/Go-BasicServer/pkg/logging"
)

type HttpClientImpl struct {
	BaseUrl    string
	HttpClient *http.Client
	Headers    *http.Header
}

type HttpClientOptions struct {
	RequestInterceptors  []middleware.HttpClientRequestInterceptor
	ResponseInterceptors []middleware.HttpClientResponseInterceptor
	BaseUrl              string
	Headers              *http.Header
}

func NewHttpClientImpl(opts *HttpClientOptions) *HttpClientImpl {

	return &HttpClientImpl{
		HttpClient: &http.Client{
			Transport: &middleware.HttpClientInterceptorManagerRoundTripper{
				Proxied:              http.DefaultTransport,
				RequestInterceptors:  opts.RequestInterceptors,
				ResponseInterceptors: opts.ResponseInterceptors,
			},
		},
		Headers: func() *http.Header {
			if opts.Headers == nil {
				return &http.Header{}
			}
			return opts.Headers
		}(),
		BaseUrl: opts.BaseUrl,
	}
}

func (c *HttpClientImpl) getFullUrl(ctx context.Context, path string) *url.URL {
	parsedURL, err := url.Parse(c.BaseUrl + path)
	if err != nil {
		logging.Error(ctx, "Error while parsing URL: "+err.Error())
		return nil
	}
	return parsedURL
}

func (c *HttpClientImpl) createHeader(body []byte, contentType ...string) http.Header {
	header := c.Headers.Clone()
	if len(contentType) > 0 {
		ct := contentType[0]
		header.Set("Content-Type", ct)
	} else {
		header.Set("Content-Type", http.DetectContentType(body))
	}
	return header
}

func (c *HttpClientImpl) Get(ctx context.Context, path string) ([]byte, error) {
	req := &http.Request{
		Method: http.MethodGet,
		URL:    c.getFullUrl(ctx, path),
		Header: *c.Headers,
	}

	response, err := c.HttpClient.Do(req)

	if err != nil {
		logging.Error(ctx, "Error while invoking url: "+path+" "+err.Error())
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		logging.Error(ctx, "Error while reading response body: "+err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		logging.Error(ctx, "Returned Error Code: "+strconv.Itoa(response.StatusCode))
		return nil, err
	}
	return body, nil
}

func (c *HttpClientImpl) Post(ctx context.Context, path string, body []byte, contentType ...string) ([]byte, error) {
	response, err := c.HttpClient.Do(&http.Request{
		Method: http.MethodPost,
		URL:    c.getFullUrl(ctx, path),
		Body:   io.NopCloser(bytes.NewReader(body)),
		Header: c.createHeader(body, contentType...),
	})

	if err != nil {
		logging.Error(ctx, "Error while invoking url: "+path+" "+err.Error())
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)

	if err != nil {
		logging.Error(ctx, "Error while reading response body: "+err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		logging.Error(ctx, "Returned Error Code: "+strconv.Itoa(response.StatusCode))
		return nil, err
	}
	return resBody, nil
}

func (c *HttpClientImpl) Patch(ctx context.Context, path string, body []byte, contentType ...string) ([]byte, error) {
	response, err := c.HttpClient.Do(&http.Request{
		Method: http.MethodPatch,
		URL:    c.getFullUrl(ctx, path),
		Body:   io.NopCloser(bytes.NewReader(body)),
		Header: c.createHeader(body, contentType...),
	})

	if err != nil {
		logging.Error(ctx, "Error while invoking url: "+path+" "+err.Error())
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)

	if err != nil {
		logging.Error(ctx, "Error while reading response body: "+err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		logging.Error(ctx, "Returned Error Code: "+strconv.Itoa(response.StatusCode))
		return nil, err
	}
	return resBody, nil
}

func (c *HttpClientImpl) Put(ctx context.Context, path string, body []byte, contentType ...string) ([]byte, error) {
	response, err := c.HttpClient.Do(&http.Request{
		Method: http.MethodPut,
		URL:    c.getFullUrl(ctx, path),
		Body:   io.NopCloser(bytes.NewReader(body)),
		Header: c.createHeader(body, contentType...),
	})

	if err != nil {
		logging.Error(ctx, "Error while invoking url: "+path+" "+err.Error())
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)

	if err != nil {
		logging.Error(ctx, "Error while reading response body: "+err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		logging.Error(ctx, "Returned Error Code: "+strconv.Itoa(response.StatusCode))
		return nil, err
	}
	return resBody, nil
}

func (c *HttpClientImpl) Delete(ctx context.Context, path string) ([]byte, error) {
	response, err := c.HttpClient.Do(&http.Request{
		Method: http.MethodDelete,
		URL:    c.getFullUrl(ctx, path),
		Header: *c.Headers,
	})

	if err != nil {
		logging.Error(ctx, "Error while invoking url: "+path+" "+err.Error())
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)

	if err != nil {
		logging.Error(ctx, "Error while reading response body: "+err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		logging.Error(ctx, "Returned Error Code: "+strconv.Itoa(response.StatusCode))
		return nil, err
	}
	return resBody, nil
}
