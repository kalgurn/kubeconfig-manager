package cmd_test

import (
	"fmt"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/cmd"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Desc",
	RunE: func(command *cobra.Command, args []string) error {
		return cmd.DelComposer(command, args)
	},
}

func TestDelete(t *testing.T) {
	output, err := executeCommand(deleteCmd, "ctx1-test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
