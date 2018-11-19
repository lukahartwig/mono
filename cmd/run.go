package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/lukahartwig/mono/client"
)

// NewRunCmd returns a cobra command that can run tasks in modules.
func NewRunCmd(cli *CLI) *cobra.Command {
	return &cobra.Command{
		Use:  "run",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			modules, err := cmd.Flags().GetStringSlice("modules")
			if err != nil {
				return err
			}
			task := args[0]
			opts := &client.RunOptions{
				Included: modules,
			}
			out, err := cli.client.Run(task, opts)
			if err != nil {
				return err
			}

			if _, err := io.Copy(os.Stdout, out); err != nil {
				return err
			}

			return nil
		},
	}
}
