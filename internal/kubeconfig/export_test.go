package kubeconfig_test

import (
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestExport(t *testing.T) {
	c := api.NewConfig()
	c1 := api.NewContext()
	u1 := api.NewAuthInfo()
	cl1 := api.NewCluster()
	c.Contexts["ctx1-test"] = c1
	c.Clusters["ctx1-test"] = cl1
	c.AuthInfos["ctx1-test"] = u1
	c.CurrentContext = "ctx1-test"

	kubeconfig.Export("ctx1-test", c)
}
