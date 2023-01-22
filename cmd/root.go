package cmd

import (
	"github.com/ruijzhan/cmds/cmd/ip"
	"github.com/ruijzhan/cmds/cmd/ss"
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "cmds",
		Short: "常用命令",
		Long:  "frequently used commands",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose")

	rootCmd.AddCommand(ss.Cmd())
	rootCmd.AddCommand(ip.Cmd())
	rootCmd.AddCommand(verCmd)

	return rootCmd
}
