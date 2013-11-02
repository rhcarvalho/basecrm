package basecrm

import (
	"net/http"
	"net/url"
)

var c httpClient = realHTTPClient{}

type httpClient interface {
	Get(url string) (*http.Response, error)
	Post(url string, data url.Values) (*http.Response, error)
}

type realHTTPClient struct{}

func (realHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
func (realHTTPClient) Post(url string, data url.Values) (*http.Response, error) {
	return http.PostForm(url, data)
}
