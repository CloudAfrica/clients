package cloudafrica

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Client  http.Client
	baseurl string
	token   string
}

func (c Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	new_path := c.baseurl + path
	req, err := http.NewRequest(method, new_path, body)
	fmt.Println("New path: " + new_path)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", "Bearer "+c.token)
	return req, nil
}

func (c Client) Get(path string) (*http.Response, error) {
	req, _ := c.newRequest("GET", path, nil)
	return c.Client.Do(req)
}

func GetClient(url string, token string) Client {
	c := http.Client{Timeout: time.Duration(5) * time.Second}
	return Client{c, url, token}
}
