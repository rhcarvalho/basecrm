package basecrm

import (
	"net/http"
	"net/url"
	"strings"
)

var httpClient *http.Client = http.DefaultClient

func NewRequest(method, url, token string, data url.Values) (req *http.Request, err error) {
	switch data {
	case nil:
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return
		}
	default:
		body := strings.NewReader(data.Encode())
		req, err = http.NewRequest(method, url, body)
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if token != "" {
		req.Header.Set("X-Pipejump-Auth", token)
		req.Header.Set("X-Futuresimple-Token", token)
	}
	return
}
