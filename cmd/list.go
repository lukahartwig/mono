package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewListCmd returns a new cobra command to list modules.
func NewListCmd(cli *CLI) *cobra.Command {
	return &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			modules, err := cli.client.List()
			if err != nil {
				return err
			}

			for _, m := range modules {
				fmt.Println(m.Name)
			}

			return nil
		},
	}
}
