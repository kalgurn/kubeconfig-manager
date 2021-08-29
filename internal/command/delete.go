package command

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

func Delete(ctx string, cfg *api.Config) error {
	if dctx, ok := cfg.Contexts[ctx]; ok {
		delete(cfg.AuthInfos, dctx.AuthInfo)
		delete(cfg.Clusters, dctx.Cluster)
		delete(cfg.Contexts, ctx)
	} else {
		return fmt.Errorf("cannot delete %s, context doesn't exist", ctx)
	}
	return nil
}
