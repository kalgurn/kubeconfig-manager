package kubeconfig

import (
	"fmt"

	"github.com/kalgurn/kubeconfig-manager/internal/utils"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Construct(ctx string, cfg *api.Config) (*api.Config, error) {
	c := api.NewConfig()
	// combine into one config json
	if ctd, err := utils.FindContextData(ctx, cfg); err != nil {
		return c, fmt.Errorf("error: %s", err)
	} else {
		if cu, err := utils.FindUser(ctd.AuthInfo, cfg); err != nil {
			return c, fmt.Errorf("error: %s", err)
		} else if cc, err := utils.FindCluster(ctd.Cluster, cfg); err != nil {
			return c, fmt.Errorf("error: %s", err)
		} else {
			c.APIVersion = cfg.APIVersion
			c.Kind = cfg.Kind
			c.CurrentContext = ctx
			c.Extensions = cfg.Extensions
			c.Preferences = cfg.Preferences
			c.Contexts[ctx] = ctd
			c.AuthInfos[ctx] = cu
			c.Clusters[ctx] = cc
		}
	}

	return c, nil
}
