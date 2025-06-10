package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "pmindt",
	Short:   "pingcode mind tools",
	Long:    "Call pmindt -h for more functions",
	Example: "pmindt -h",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			// 当初没有参数时 提示用户
			tip()
		}
	},
}

func tip() {
	fmt.Println(`You can try using -h to get more information`)
}

func init() {
}

func ExecCmd() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
