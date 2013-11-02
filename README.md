# Base CRM API wrapper for Go

This is an unofficial wrapper in Go for accessing the [Base CRM API](http://dev.futuresimple.com/api/).

## Get it!

After you've setup your Go environment, it is as easy as:

    go get github.com/rhcarvalho/basecrm

## Use it!

Just import `github.com/rhcarvalho/basecrm` into your code and you're good to go.
Take a look at this example:

```go
package main

import (
	"fmt"
	"os"
	// ...
	"github.com/rhcarvalho/basecrm"
)

func main() {
	email, password := os.Getenv("BASECRM_EMAIL"), os.Getenv("BASECRM_PASSWORD")
	s, err := basecrm.NewSession(email, password)
	switch err {
	case nil:
		fmt.Printf("Session TOKEN: %s\n", s.Token)
	case basecrm.NotAuthenticated:
		fmt.Printf("Failed to authenticate.")
		return
	default:
		fmt.Printf("Error: %v", err)
		return
	}
	account, err := s.Account()
	switch err {
	case nil:
		fmt.Printf("Account: %v\n", account)
	default:
		fmt.Printf("Error: %v", err)
		return
	}
}
```
