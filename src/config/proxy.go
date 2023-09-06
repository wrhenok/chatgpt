package config

import (
	"net/http"
)

type Client struct {
	config ClientConfig
}

const (
	apiURLv1                       = "https://api2.ycdl.pro/v1"
	defaultEmptyMessagesLimit uint = 300
)

type ClientConfig struct {
	authToken string

	HTTPClient *http.Client

	BaseURL string
	OrgID   string

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		HTTPClient: http.DefaultClient,
		BaseURL:    apiURLv1,
		OrgID:      "",
		authToken:  authToken,

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{config}
}
