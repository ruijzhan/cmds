package main

import (
	"github.com/ruijzhan/cmds/cmd"
)

func main() {

	rootCmd := cmd.NewRootCmd()

	rootCmd.Execute()
}
