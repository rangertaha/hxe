package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client represents an API client for interacting with the hxe API
type Client struct {
	GenericClient
}

type GenericClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new API client
func New(baseURL string) *Client {
	return &Client{
		GenericClient: GenericClient{
			baseURL: baseURL,
			httpClient: &http.Client{
				Timeout: time.Second * 30,
			},
		},
	}
}

// NewClient creates a new API client
func NewGenericClient(baseURL string) *GenericClient {
	return &GenericClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

// Generic HTTP methods

func (c *GenericClient) Get(path string, v interface{}) error {
	return c.DoRequest(http.MethodGet, path, nil, v)
}

func (c *GenericClient) Post(path string, body interface{}, v interface{}) error {
	return c.DoRequest(http.MethodPost, path, body, v)
}

func (c *GenericClient) Put(path string, body interface{}, v interface{}) error {
	return c.DoRequest(http.MethodPut, path, body, v)
}

func (c *GenericClient) Delete(path string) error {
	return c.DoRequest(http.MethodDelete, path, nil, nil)
}

func (c *GenericClient) DoRequest(method, path string, body interface{}, v interface{}) error {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, c.baseURL+path, &buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errResp struct {
			Error string `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		return fmt.Errorf("API error: %s", errResp.Error)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
