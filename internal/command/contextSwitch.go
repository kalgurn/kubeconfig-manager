package command

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

func ContextSwitch(ctx string, cfg *api.Config) error {
	found := false
	for k := range cfg.Contexts {
		if k == ctx {
			found = true
		}
	}
	if found {
		cfg.CurrentContext = ctx
	} else {
		return fmt.Errorf("no context with name %s in the current config", ctx)
	}
	return nil
}
