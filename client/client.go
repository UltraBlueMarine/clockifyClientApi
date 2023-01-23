package client

import (
	"errors"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be 0")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &Middleware.LoggingRoundTripper{
				Logger: os.Stdout,
				Next:   http.DefaultClient,
			},
		},
	}, nil
}
