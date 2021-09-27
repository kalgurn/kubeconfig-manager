package kubeconfig

import (
	"fmt"

	"facette.io/natsort"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Contexts(cfg *api.Config) []string {
	var contexts []string
	for k := range cfg.Contexts {
		contexts = append(contexts, k)
	}
	return contexts
}

func SortedContexts(contexts []string) []string {
	natsort.Sort(contexts)
	return contexts
}

func CurrentContext(cfg *api.Config) string {
	return cfg.CurrentContext
}

func ContextExists(ctx string, cfg *api.Config) error {
	for _, k := range Contexts(cfg) {
		if k == ctx {
			return nil
		}
	}
	return fmt.Errorf("no context exists: %s", ctx)
}
