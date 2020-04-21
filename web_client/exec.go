package web_client

import (
	"fmt"
	"net/http"
)

func DoOps(client *http.Client) error {
	resp, err := client.Get("http://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func UseDefaultClient() error {
	resp, err := http.Get("https://www.golang.org")
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

// Controller embeds (inherit) http.Clint and uses it internally

type Controller struct {
	*http.Client
}

func (c *Controller) DoOps() error {
	resp, err := c.Client.Get("http://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}
