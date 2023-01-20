package cmd

import (
	"fmt"
	"os"

	"github.com/ruijzhan/cmds/cmd/ip"
	"github.com/ruijzhan/cmds/cmd/ss"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ss.Cmd())
	rootCmd.AddCommand(ip.Cmd())
}

var rootCmd = &cobra.Command{
	Use:   "cmds",
	Short: "常用命令",
	Long:  "frequently used commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprint(os.Stdout, "haha\n")
	},
}

func Execute() {
	rootCmd.Execute()
}
