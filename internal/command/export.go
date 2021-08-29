package command

import (
	"fmt"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"github.com/kalgurn/kubeconfig-manager/internal/utils"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Export(ctx string, cfg *api.Config) error {
	if c, err := utils.ConstructConfig(ctx, cfg); err != nil {
		return fmt.Errorf("error: %s", err)
	} else {
		kubeconfig.Export(ctx, c)
	}
	return nil
}
