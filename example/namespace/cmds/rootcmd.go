package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "ns-practice-tool",
	Short: "tool for namespace learning",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("learning namespace...")
	},
}

func init() {
	RootCmd.AddCommand(runCommand, execCommand)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
