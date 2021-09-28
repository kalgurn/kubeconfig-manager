package kubeconfig_test

import (
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestMerge(t *testing.T) {
	config1 := api.NewConfig()
	config2 := api.NewConfig()

	kubeconfig.Merge(config1, config2)

}
