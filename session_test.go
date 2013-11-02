package basecrm

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
)

type fakeTransport struct {
	Answer     func(*http.Request) string
	StatusCode int
	requests   []*http.Request
	sync.Mutex
}

func (ft *fakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ft.Lock()
	ft.requests = append(ft.requests, req)
	ft.Unlock()
	res := &http.Response{
		Status:     fmt.Sprintf("%d %s", ft.StatusCode, http.StatusText(ft.StatusCode)),
		StatusCode: ft.StatusCode,
		Header:     make(http.Header),
		Body:       ioutil.NopCloser(strings.NewReader(ft.Answer(req))),
	}
	return res, nil
}

func TestAuthenticationSuccess(t *testing.T) {
	ft := &fakeTransport{
		Answer: func(req *http.Request) string {
			return fmt.Sprintf(`{
				"authentication": {
					"token": "TOKEN_%s"
				}
			}`, req.FormValue("email"))
		},
		StatusCode: http.StatusOK,
	}
	httpClient.Transport = ft
	email, password := "user@company.com", "secret_password"
	s := NewSession(email, password)
	req := ft.requests[0]
	if method := req.Method; method != "POST" {
		t.Errorf("wrong method: %s", method)
	}
	if url := req.URL; url.String() != AuthenticationEndpoint {
		t.Errorf("wrong url: %s", url)
	}
	if token := s.Token; token != "TOKEN_user@company.com" {
		t.Errorf("bad token: %s", token)
	}
}

func TestAuthenticationFailure(t *testing.T) {
	ft := &fakeTransport{
		Answer: func(req *http.Request) string {
			return `{
				"authentication": {
				}
			}`
		},
		StatusCode: http.StatusUnauthorized,
	}
	httpClient.Transport = ft
	email, password := "user@company.com", "secret_password"
	s := NewSession(email, password)
	req := ft.requests[0]
	if method := req.Method; method != "POST" {
		t.Errorf("wrong method: %s", method)
	}
	if url := req.URL; url.String() != AuthenticationEndpoint {
		t.Errorf("wrong url: %s", url)
	}
	if token := s.Token; token != "" {
		t.Errorf("bad token: %s", token)
	}
}
