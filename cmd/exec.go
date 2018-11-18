package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

func NewExecCmd(cli *CLI) *cobra.Command {
	return &cobra.Command{
		Use:  "exec",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				out io.Reader
				err error
			)

			if len(args) == 1 {
				out, err = cli.client.Exec(args[0])
			} else {
				out, err = cli.client.Exec(args[0], args[1:]...)
			}

			_, err = io.Copy(os.Stdout, out)
			return err
		},
	}
}
