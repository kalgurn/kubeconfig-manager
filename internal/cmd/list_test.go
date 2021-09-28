package cmd_test

import (
	"fmt"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/cmd"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Desc",
	RunE: func(command *cobra.Command, args []string) error {
		return cmd.ListComposer(command, args)
	},
}

func TestList(t *testing.T) {
	output, err := executeCommand(listCmd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
