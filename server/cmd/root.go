package cmd

import (
	"github.com/kmiit/vivi/cmd/flags"
	"github.com/kmiit/vivi/utils/log"
	"github.com/spf13/cobra"
)

const TAG = "root"

const logDesc = `Log level:
	0: Fatal
	1: Error
	2: Warn
	3: Info
	4: Debug
	5: Verbose`

var rootCmd = &cobra.Command{
	Use:   "vivi",
	Short: "A small file server written in go",
	Long:  `A file server backend written in golang`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.F(TAG, err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flags.ConfigFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().IntVarP(&flags.LogLevel, "loglevel", "l", 3, logDesc)
}
