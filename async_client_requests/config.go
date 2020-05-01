package async_client_requests

import (
	"net/http"
)

// NewClient, creates a client struct with response and error channel

type Client struct {
	*http.Client
	Resp chan *http.Response
	Err  chan error
}

//perform a Get then returns response/error to channels
func (c *Client) AsyncGet(url string) {
	resp, err := c.Get(url)
	if err != nil {
		c.Err <- err
		return
	}
	c.Resp <- resp
}

func NewClient(client *http.Client, bufferSize int) *Client {
	respChan := make(chan *http.Response, bufferSize)
	errChan := make(chan error, bufferSize)
	return &Client{
		Client: client,
		Resp:   respChan,
		Err:    errChan,
	}
}
