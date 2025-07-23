package http

import (
	"fmt"
	"net/http"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Get(url string) (*http.Response, error) {
	const op = "op.http.Get"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	
	return resp, nil
}
