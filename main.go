package main

import (
	"github.com/lukahartwig/mono/cmd"
)

func main() {
	cli := cmd.New()

	_ = cmd.NewRootCmd(cli).Execute()
}
