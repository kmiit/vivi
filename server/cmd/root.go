/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kmiit/vivi/cmd/flags"
	"github.com/kmiit/vivi/utils/log"
	"github.com/spf13/cobra"
)

const TAG = "root"

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
}
