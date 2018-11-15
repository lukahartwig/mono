package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/lukahartwig/mono/client"
)

var (
	c client.Client

	configFile string
)

func init() {
	cobra.OnInitialize(initClient)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", ".mono.yml", "config file")
}

var rootCmd = &cobra.Command{
	Use: "mono",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func initClient() {
	initConfig()

	opts := &client.Options{
		Root: viper.GetString("root"),
	}

	c = client.New(opts)
}

func initConfig() {
	viper.SetConfigFile(configFile)
	viper.SetDefault("root", ".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("cannot read config:", err)
	}
}
