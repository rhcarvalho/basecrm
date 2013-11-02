package basecrm

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
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
