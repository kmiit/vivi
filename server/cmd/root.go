/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/kmiit/vivi/cmd/flags"
	
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vivi",
	Short: "A small file server written in go",
	Long: `A file server backend written in golang`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&flags.ConfigFile, "config", "", "config file")
}


