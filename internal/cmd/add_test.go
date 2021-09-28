package cmd_test

import (
	"fmt"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/cmd"
	"github.com/spf13/cobra"
)

func TestAdd(t *testing.T) {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Desc",
		RunE: func(command *cobra.Command, args []string) error {
			return cmd.AddComposer(command, args)
		},
	}
	output, err := executeCommand(addCmd, "ctx1-test.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
