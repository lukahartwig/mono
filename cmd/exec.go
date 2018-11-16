package cmd

import (
	"github.com/spf13/cobra"
)

func NewExecCmd(cli *CLI) *cobra.Command {
	return &cobra.Command{
		Use:  "exec",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				return cli.client.Exec(args[0])
			} else {
				return cli.client.Exec(args[0], args[1:]...)
			}
		},
	}
}
