package client

import (
	"bytes"
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

func (c Client) GetClients() ([]logic.ClientData, error) {
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

	var clientsData []logic.ClientData
	if err = json.Unmarshal(body, &clientsData); err != nil {
		return nil, err
	}

	return clientsData, nil
}

func (c Client) GetClientById(id string) (logic.ClientData, error) {
	req, err := http.NewRequest("GET", "https://api.clockify.me/api/v1/workspaces/"+c.workspaceID+"/clients/"+id, nil)
	if err != nil {
		return logic.ClientData{}, err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return logic.ClientData{}, err
	}

	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return logic.ClientData{}, err
	}

	var clientData logic.ClientData
	if err = json.Unmarshal(body, &clientData); err != nil {
		return logic.ClientData{}, err
	}
	return clientData, nil
}

func (c Client) InsertNewClient(name string, note string) error {
	newUser := logic.InsertedData{Name: name, Note: note}
	body, err := json.Marshal(newUser)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://api.clockify.me/api/v1/workspaces/"+c.workspaceID+"/clients", bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)
	return nil
}

func (c Client) UpdateClient(id string, newNote string, archived bool) error {
	client, err := c.GetClientById(id)
	if err != nil {
		return err
	}

	newUser := logic.UpdatedData{Archived: false, Name: client.Name, Note: newNote}
	body, err := json.Marshal(newUser)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", "https://api.clockify.me/api/v1/workspaces/"+c.workspaceID+"/clients/"+client.ID, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)
	return nil
}

func (c Client) DeleteClient(id string) error {
	client, err := c.GetClientById(id)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", "https://api.clockify.me/api/v1/workspaces/"+c.workspaceID+"/clients/"+client.ID, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)
	return nil
}
