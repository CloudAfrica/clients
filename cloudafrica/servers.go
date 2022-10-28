package cloudafrica

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ServersAPI struct {
	client *Client
}

type Server struct {
	Name   string `json:"server/name"`
	CPUs   int    `json:"server/cpus"`
	RamMiB int    `json:"server/ram-mib"`
	//TODO: disks
	//TODO: networks
}

type Servers struct {
	Servers []Server
}

func (c ServersAPI) List() ([]Server, error) {
	resp, err := c.client.Get("/servers")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body: %s\n", body)
	if err != nil {
		return nil, err
	}
	var servers Servers
	json.Unmarshal(body, &servers)
	return servers.Servers, nil
}
