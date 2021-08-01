package utils

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

func ConstructConfig(ctx string, cfg *api.Config) (*api.Config, error) {
	c := api.NewConfig()
	// combine into one config json
	if ctd, err := FindContextData(ctx, cfg); err != nil {
		return c, fmt.Errorf("Error: %s", err)
	} else {
		if cu, err := FindUser(ctd.AuthInfo, cfg); err != nil {
			return c, fmt.Errorf("Error: %s", err)
		} else if cc, err := FindCluster(ctd.Cluster, cfg); err != nil {
			return c, fmt.Errorf("Error: %s", err)
		} else {
			c.APIVersion = cfg.APIVersion
			c.Kind = cfg.Kind
			c.CurrentContext = cfg.CurrentContext
			c.Extensions = cfg.Extensions
			c.Preferences = cfg.Preferences
			c.Contexts[ctx] = ctd
			c.AuthInfos[ctx] = cu
			c.Clusters[ctx] = cc
			fmt.Println(c)
		}
	}

	return c, nil
}
