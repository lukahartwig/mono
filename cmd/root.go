package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/lukahartwig/mono/client"
	"github.com/lukahartwig/mono/module"
)

// CLI is the context that is passed to all the commands.
type CLI struct {
	client client.Client
}

// New returns a new CLI instance
func New() *CLI {
	return &CLI{}
}

// NewRootCmd returns the root cobra command
func NewRootCmd(cli *CLI) *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "mono",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			configFile, err := cmd.Flags().GetString("config")
			if err != nil {
				return err
			}
			config, err := loadConfig(configFile)
			if err != nil {
				return err
			}
			resolver := module.NewResolver(config.Root)
			cli.client = client.New(resolver)
			return nil
		},
		TraverseChildren: true,
	}

	rootCmd.Flags().StringP("config", "c", ".mono.yml", "config file")

	rootCmd.AddCommand(
		NewExecCmd(cli),
		NewListCmd(cli),
		NewRunCmd(cli),
	)

	return rootCmd
}

type config struct {
	Root string
}

func loadConfig(configFile string) (config, error) {
	viper.SetConfigFile(configFile)
	viper.SetDefault("root", ".")

	if err := viper.ReadInConfig(); err != nil {
		return config{}, err
	}

	var c config
	if err := viper.Unmarshal(&c); err != nil {
		return config{}, err
	}

	return c, nil
}
