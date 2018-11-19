package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/lukahartwig/mono/client"
)

// NewExecCmd returns a new cobra command to run commands in modules.
func NewExecCmd(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "exec",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			modules, err := cmd.Flags().GetStringSlice("modules")
			if err != nil {
				return err
			}
			opts := &client.ExecOptions{
				Included: modules,
			}
			out, err := cli.client.Exec(args[0], args[1:], opts)
			if err != nil {
				return err
			}
			_, err = io.Copy(os.Stdout, out)
			return err
		},
	}

	cmd.Flags().StringSliceP("modules", "m", nil, "apply command to these modules")

	return cmd
}
