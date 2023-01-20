package ss

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ssCmd = &cobra.Command{
	Use:   "ss",
	Short: "start/stop shadowsocks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("还没实现")
	},
}

func Cmd() *cobra.Command {
	return ssCmd
}
