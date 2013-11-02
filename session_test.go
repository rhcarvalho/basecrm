package basecrm

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"testing"
)

type fakeHTTPClient struct {
	Answer     func(*http.Request) string
	StatusCode int
	requests   []*http.Request
	sync.Mutex
}

func (fc *fakeHTTPClient) RoundTrip(req *http.Request) (*http.Response, error) {
	fc.Lock()
	fc.requests = append(fc.requests, req)
	fc.Unlock()
	res := &http.Response{
		Status:     fmt.Sprintf("%d %s", fc.StatusCode, http.StatusText(fc.StatusCode)),
		StatusCode: fc.StatusCode,
		Header:     make(http.Header),
		Body:       ioutil.NopCloser(strings.NewReader(fc.Answer(req))),
	}
	return res, nil
}

func (fc *fakeHTTPClient) Post(url string, data url.Values) (*http.Response, error) {
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return fc.RoundTrip(req)
}

func TestAuthentication(t *testing.T) {
	fc := &fakeHTTPClient{
		Answer: func(req *http.Request) string {
			return fmt.Sprintf(`{
				"authentication": {
					"token": "TOKEN_%s"
				}
			}`, req.FormValue("email"))
		},
		StatusCode: http.StatusOK,
	}
	c = fc
	email, password := "user@company.com", "secret_password"
	s := NewSession(email, password)
	req := fc.requests[0]
	if method := req.Method; method != "POST" {
		t.Errorf("wrong method: %s", method)
	}
	if url := req.URL; url.String() != "https://sales.futuresimple.com/api/v1/authentication.json" {
		t.Errorf("wrong url: %s", url)
	}
	if token := s.Token; token != "TOKEN_user@company.com" {
		t.Errorf("bad token: %s", token)
	}
}
