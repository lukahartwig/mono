package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		modules, err := c.List()
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range modules {
			fmt.Println(m.Name)
		}
	},
}
