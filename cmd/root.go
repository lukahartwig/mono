package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/lukahartwig/mono/client"
	"github.com/lukahartwig/mono/module"
)

var (
	configFile string
)

type CLI struct {
	client client.Client
}

func New() *CLI {
	return &CLI{}
}

func NewRootCmd(cli *CLI) *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "mono",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := loadConfig()
			if err != nil {
				return err
			}

			resolver := module.NewResolver(config.Root)

			opts := &client.Options{}
			cli.client = client.New(resolver, opts)

			return nil
		},
		TraverseChildren: true,
	}

	rootCmd.Flags().StringVarP(&configFile, "config", "c", ".mono.yml", "config file")

	rootCmd.AddCommand(
		NewExecCmd(cli),
		NewListCmd(cli),
	)

	return rootCmd
}

type config struct {
	Root string
}

func loadConfig() (config, error) {
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
