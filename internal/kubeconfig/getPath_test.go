package kubeconfig_test

import (
	"fmt"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/kubeconfig"
)

func TestGetConfigPath(t *testing.T) {
	makeConfig()
	fmt.Println("getPath-test")
	kubeconfig.GetConfigPath()
}
