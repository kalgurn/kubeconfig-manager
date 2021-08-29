package kubeconfig_test

import (
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestLoad(t *testing.T) {
	c := api.NewConfig()
	kubeconfig.Export("cfg-test", c)

	kubeconfig.Load("cfg-test.yaml")
	kubeconfig.Load("test.yaml")

}
