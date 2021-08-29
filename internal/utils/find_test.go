package utils_test

import (
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/utils"
	"k8s.io/client-go/tools/clientcmd/api"
)

var (
	c   = api.NewConfig()
	c1  = api.NewContext()
	u1  = api.NewAuthInfo()
	cl1 = api.NewCluster()
)

func TestFindUser(t *testing.T) {

	c.AuthInfos["ctx1-test"] = u1
	utils.FindUser("ctx1-test", c)
	utils.FindUser("ctx2-test", c)

}

func TestFindContextData(t *testing.T) {
	c.Contexts["ctx1-test"] = c1
	utils.FindContextData("ctx1-test", c)
	utils.FindContextData("ctx2-test", c)
}

func TestFindCluster(t *testing.T) {
	c.Clusters["ctx1-test"] = cl1
	utils.FindCluster("ctx1-test", c)
	utils.FindCluster("ctx2-test", c)
}
