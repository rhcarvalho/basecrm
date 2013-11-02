package basecrm

import (
	"encoding/json"
	"net/url"
)

type Session struct {
	Token string
}

type jsonResponse struct {
	Authentication struct {
		Token string `json:"token"`
	} `json:"authentication"`
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
	auth := &jsonResponse{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(auth)
	if err != nil {
		panic(err)
	}
	return &Session{
		Token: auth.Authentication.Token,
	}
}
