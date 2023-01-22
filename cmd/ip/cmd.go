package ip

import (
	"github.com/spf13/cobra"
)

var s string

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "view my network's external IP",
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func Cmd() *cobra.Command {
	return ipCmd
}

func init() {
	ipCmd.Flags().StringVarP(&s, "site", "s", "", "site to ask for my IP")
}
