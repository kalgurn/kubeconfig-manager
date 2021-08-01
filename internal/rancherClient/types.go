package rancherClient

type clusters struct {
	Data []data `json:"data"`
}
type data struct {
	Id string `json:"id"`
}

type config struct {
	Config string `json:"config"`
}
