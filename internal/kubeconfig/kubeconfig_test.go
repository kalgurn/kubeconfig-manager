package kubeconfig_test

import (
	"os"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"k8s.io/client-go/tools/clientcmd/api"
)

func makeConfig() *api.Config {
	cfg := api.NewConfig()
	cfg.Contexts["ctx1-test"] = api.NewContext()
	cfg.Clusters["ctx1-test"] = api.NewCluster()
	cfg.AuthInfos["ctx1-test"] = api.NewAuthInfo()
	cfg.CurrentContext = "ctx1-test"
	kubeconfig.Export("ctx1-test", cfg)
	os.Setenv("KUBECONFIG", "ctx1-test.yaml")
	return cfg
}
