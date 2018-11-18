package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

// NewRunCmd returns a cobra command that can run tasks in modules
func NewRunCmd(cli *CLI) *cobra.Command {
	return &cobra.Command{
		Use:  "run",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			task := args[0]

			out, err := cli.client.RunTask(task)
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
