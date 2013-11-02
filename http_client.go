package basecrm

import (
	"net/http"
	"net/url"
)

var c httpClient = realHTTPClient{}

type httpClient interface {
	Post(url string, data url.Values) (*http.Response, error)
}

type realHTTPClient struct{}

func (realHTTPClient) Post(url string, data url.Values) (*http.Response, error) {
	return http.PostForm(url, data)
}
