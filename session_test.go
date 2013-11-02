package basecrm

import (
	"fmt"
	"net/http"
	"testing"
)

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
	s, err := NewSession(email, password)
	if err != nil {
		t.Fatalf("session error: %s", err)
	}
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
	_, err := NewSession(email, password)
	if err != NotAuthenticated {
		t.Errorf("should be authentication error: %v", err)
	}
	req := ft.requests[0]
	if method := req.Method; method != "POST" {
		t.Errorf("wrong method: %s", method)
	}
	if url := req.URL; url.String() != AuthenticationEndpoint {
		t.Errorf("wrong url: %s", url)
	}
}
