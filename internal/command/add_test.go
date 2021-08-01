package command

import (
	"testing"

	"k8s.io/client-go/tools/clientcmd/api"
)

func TestAdd(t *testing.T) {
	config1 := api.NewConfig()
	config2 := api.NewConfig()

	Add(config1, config2)

}
