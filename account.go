package basecrm

import (
	"encoding/json"
	"net/http"
)

type Account struct {
	Name         string `json:"name"`
	Id           int    `json:"id"`
	Timezone     string `json:"timezone"`
	CurrencyName string `json:"currency_name"`
}

func (s *Session) Account() (*Account, error) {
	resp, err := c.Get(AccountEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, NotAuthenticated
	}
	message := &struct {
		Account *Account
	}{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(message)
	if err != nil {
		return nil, err
	}
	return message.Account, nil
}
