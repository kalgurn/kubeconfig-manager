package rancherClient

type Clusters struct {
	Data []Data `json:"data"`
}
type Data struct {
	Id string `json:"id"`
}

type Config struct {
	Config string `json:"config"`
}
