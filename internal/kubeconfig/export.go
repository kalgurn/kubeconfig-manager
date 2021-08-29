package kubeconfig

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

func Export(ctx string, cfg *api.Config) {
	configPath := fmt.Sprintf("%s.yaml", ctx)
	Save(configPath, cfg)
}
