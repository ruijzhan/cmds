package cmd

import (
	"fmt"
	"os"

	"github.com/ruijzhan/cmds/constant"

	"github.com/spf13/cobra"
)

var verCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "version: %s, commit: %s\n", constant.Version, constant.Commit)
	},
}
