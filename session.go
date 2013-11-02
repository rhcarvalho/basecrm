package basecrm

import (
	"encoding/json"
	"errors"
	"net/url"
)

var NotAuthenticated = errors.New("not authenticated")

type Session struct {
	Token string `json:"token"`
}

func NewSession(email, password string) *Session {
	resp, err := c.Post("https://sales.futuresimple.com/api/v1/authentication.json", url.Values{
		"email":    {email},
		"password": {password},
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	message := &struct {
		Authentication *Session
	}{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(message)
	if err != nil {
		panic(err)
	}
	return message.Authentication
}
