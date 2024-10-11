package cmd

import (
  "fmt"

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
    fmt.Println("Vivi version: 0.0.1")
  },
}