package command

import (
	"testing"

	"k8s.io/client-go/tools/clientcmd/api"
)

func TestContextSwitch(t *testing.T) {
	c := api.NewConfig()
	c1 := api.NewContext()
	c2 := api.NewContext()
	c.Contexts["ctx1"] = c1
	c.Contexts["ctx2"] = c2
	c.CurrentContext = "ctx1"
	ContextSwitch("ctx2", c)
	ContextSwitch("ctx3", c)

}
