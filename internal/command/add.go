package command

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

func Add(importCfg *api.Config, cfg *api.Config) error {
	for cl, v := range importCfg.Clusters {
		cfg.Clusters[cl] = v
	}
	for u, v := range importCfg.AuthInfos {
		cfg.AuthInfos[u] = v
	}
	for c, v := range importCfg.Contexts {
		cfg.Contexts[c] = v
		fmt.Println("added data for context", c)
	}
	return nil
}
