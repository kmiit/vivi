package cmd

import (
	"fmt"

	"github.com/kmiit/vivi/utils/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of vivi",
	Long:  `Show the current version of vivi which you are executing.`,
	Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`Build Date: %s
Go Version: %s
Version: %s
`,
	version.BuildDate,
	version.GoVersion,
	version.Version)
	},
}
