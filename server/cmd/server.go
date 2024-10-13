package cmd

import (
    "context"
    "fmt"
    "sync"

	//"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/cmd/flags"
	"github.com/kmiit/vivi/utils/config"
	"github.com/kmiit/vivi/utils/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the vivi server",
	Long:  `Run the vivi server in frontend.`,
  	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	var wg sync.WaitGroup
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunServer(config.Parse(flags.ConfigFile))
	}()
	
	wg.Wait()
	fmt.Println("Done")
}