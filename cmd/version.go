package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version = "1.2.4-DEV"
	Commit  = ""
	Date    = ""
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", Version)
		fmt.Println("Date:", Date)
		fmt.Println("Commit:", Commit)
	},
}
