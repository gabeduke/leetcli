package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "leetcli",
	Short: "leetcli is the master cli",
	Long: `CLI to interact will all of the leetsphere`,
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	},
  }

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}