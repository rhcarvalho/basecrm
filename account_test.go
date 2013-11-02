package basecrm

import (
	"net/http"
	"testing"
)

func TestGetAccountAuthorized(t *testing.T) {
	fc := &fakeHTTPClient{
		Answer: func(req *http.Request) string {
			return `{
				"account": {
					"name": "myaccount",
					"id": 60,
					"timezone": "UTC",
					"currency_name": "US Dollar"
				}
			}`
		},
		StatusCode: http.StatusOK,
	}
	c = fc
	s := &Session{"TOKEN_account_test"}
	account, err := s.Account()
	if err != nil {
		t.Fatalf("account error: %s", err)
	}
	req := fc.requests[0]
	if method := req.Method; method != "GET" {
		t.Errorf("wrong method: %s", method)
	}
	if url := req.URL; url.String() != "https://sales.futuresimple.com/api/v1/account.json" {
		t.Errorf("wrong url: %s", url)
	}
	tokenPipejump := req.Header.Get("X-Pipejump-Auth")
	tokenFutureSimple := req.Header.Get("X-Futuresimple-Token")
	if !(tokenPipejump == s.Token && tokenFutureSimple == s.Token) {
		t.Errorf("missing token: %s, %s", tokenPipejump, tokenFutureSimple)
	}
	if name := account.Name; name != "myaccount" {
		t.Errorf("wrong account name: %s", name)
	}
	if id := account.Id; id != 60 {
		t.Errorf("wrong account id: %d", id)
	}
	if tz := account.Timezone; tz != "UTC" {
		t.Errorf("wrong account timezone: %s", tz)
	}
	if currency := account.CurrencyName; currency != "US Dollar" {
		t.Errorf("wrong account currency: %s", currency)
	}
}

func TestGetAccountUnauthorized(t *testing.T) {
	fc := &fakeHTTPClient{
		Answer: func(req *http.Request) string {
			return `{"message":"Not authenticated"}`
		},
		StatusCode: http.StatusUnauthorized,
	}
	c = fc
	s := &Session{"TOKEN_account_test"}
	account, err := s.Account()
	if err != NotAuthenticated {
		t.Errorf("should be authentication error: %v", err)
	}
	if account != nil {
		t.Errorf("account should be nil: %v", account)
	}
	req := fc.requests[0]
	if method := req.Method; method != "GET" {
		t.Errorf("wrong method: %s", method)
	}
	if url := req.URL; url.String() != "https://sales.futuresimple.com/api/v1/account.json" {
		t.Errorf("wrong url: %s", url)
	}
}
