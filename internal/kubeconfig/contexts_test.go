package kubeconfig_test

import (
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestContexts(t *testing.T) {
	cfg := api.NewConfig()
	cfg.Contexts["ctx1-test"] = api.NewContext()
	kubeconfig.Contexts(cfg)
}
func TestSortedContexts(t *testing.T) {
	cfg := api.NewConfig()
	cfg.Contexts["ctx1-test"] = api.NewContext()
	contexts := kubeconfig.Contexts(cfg)
	kubeconfig.SortedContexts(contexts)
}
func TestContextExist(t *testing.T) {
	cfg := api.NewConfig()
	cfg.Contexts["ctx1-test"] = api.NewContext()
	kubeconfig.ContextExists("ctx1-test", cfg)
}
func TestContextExistFail(t *testing.T) {
	cfg := api.NewConfig()
	cfg.Contexts["ctx1-test"] = api.NewContext()
	kubeconfig.ContextExists("ctx2-test", cfg)
}
func TestCurrentContext(t *testing.T) {
	cfg := api.NewConfig()
	cfg.Contexts["ctx1-test"] = api.NewContext()
	cfg.CurrentContext = "ctx1-test"
	kubeconfig.CurrentContext(cfg)
}
