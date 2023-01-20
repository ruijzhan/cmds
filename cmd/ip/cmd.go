package ip

import (
	"github.com/spf13/cobra"
)

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
