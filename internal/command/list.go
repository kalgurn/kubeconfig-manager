package command

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

func List(cfg *api.Config) {
	for k := range cfg.Contexts {
		if cfg.CurrentContext == k {
			fmt.Println(k, "<-")
		} else {
			fmt.Println(k)
		}

	}
}
