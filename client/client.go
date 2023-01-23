package client

import (
	"clockifyClientApi/logic"
	"clockifyClientApi/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client      *http.Client
	apiKey      string
	workspaceID string
}

func NewClient(timeout time.Duration, _apiKey string, _worspaceID string) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be 0")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &middleware.LoggingRoundTripper{
				Logger: os.Stdout,
				Next:   http.DefaultTransport,
			},
		},
		apiKey:      _apiKey,
		workspaceID: _worspaceID,
	}, nil
}

func (c Client) GetClients() ([]logic.AssetsResponses, error) {
	req, err := http.NewRequest("GET", "https://api.clockify.me/api/v1/workspaces/"+c.workspaceID+"/clients", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r logic.AssetsResponses
	if err = json.Unmarshal(body, &r.Assets); err != nil {
		return nil, err
	}

	return r, nil
}
