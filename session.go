package basecrm

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

var NotAuthenticated = errors.New("not authenticated")

type Session struct {
	Token string `json:"token"`
}

func NewSession(email, password string) (*Session, error) {
	req, err := NewRequest("POST", AuthenticationEndpoint, "", url.Values{
		"email":    {email},
		"password": {password},
	})
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, NotAuthenticated
	}
	message := &struct {
		Authentication *Session
	}{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(message)
	if err != nil {
		return nil, err
	}
	return message.Authentication, nil
}
