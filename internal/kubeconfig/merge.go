package kubeconfig

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

func Merge(importCfg *api.Config, cfg *api.Config) error {
	if len(importCfg.Contexts) >= 1 {
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

	return fmt.Errorf("no config to import")
}
