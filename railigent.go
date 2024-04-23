package railigentx

import (
	"encoding/base64"
	"net/http"
	"time"
)

// Client holds the configuration for the Railigent X API client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	AuthHeader string
}

// NewClient creates a new Railigent X API client with default settings
func NewClient(baseURL, username, password string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
		AuthHeader: getBasicAuthHeader(username, password),
	}
}

func getBasicAuthHeader(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
