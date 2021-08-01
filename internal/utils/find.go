package utils

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

func FindContextData(ctx string, cfg *api.Config) (*api.Context, error) {

	for k, v := range cfg.Contexts {
		if k == ctx {
			return v, nil
		}
	}
	return &api.Context{}, fmt.Errorf("No context called %s found in config", ctx)
}

func FindCluster(c string, cfg *api.Config) (*api.Cluster, error) {
	for k, v := range cfg.Clusters {
		if k == c {
			return v, nil
		}
	}
	return &api.Cluster{}, fmt.Errorf("No cluster with name %s found in config", c)
}

func FindUser(u string, cfg *api.Config) (*api.AuthInfo, error) {
	for k, v := range cfg.AuthInfos {
		if k == u {
			return v, nil
		}
	}
	return &api.AuthInfo{}, fmt.Errorf("No authInfo for name %s found in config", u)
}
