package async_client_requests

import "time"

//execute a list of urls
func FetchAll(urls []string, c *Client) {
	for _, url := range urls {
		go c.AsyncGet(url)
		time.Sleep(300 * time.Millisecond)
	}
}
