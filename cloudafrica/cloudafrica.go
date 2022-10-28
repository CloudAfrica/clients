package cloudafrica

import ()

type CloudAfrica struct {
	Servers ServersAPI
}

func Init(url string, token string) CloudAfrica {
	client := GetClient(url, token)
	servers := ServersAPI{&client}
	return CloudAfrica{servers}
}
