package request

import (
	"fmt"
	"net/http"
)

type Request struct {
	URL string
}

func NewRequest(url string) *Request {
	return &Request{
		URL: url,
	}
}

func (r *Request) GetResponse() (*http.Response, error) {
	const op = "op.request.GetResponse"
	res, err := http.Get(r.URL)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return res, nil
}
