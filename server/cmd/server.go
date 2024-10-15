package cmd

import (
	"context"
	"sync"

	"github.com/kmiit/vivi/utils/config"
	"github.com/kmiit/vivi/utils/db"
	"github.com/kmiit/vivi/utils/log"
	"github.com/kmiit/vivi/utils/server"
	"github.com/kmiit/vivi/utils/storage"

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
		config.InitConfig()
		storage.InitStorage()
		db.InitDatabase()

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
		server.InitServer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		storage.WatchStorage()
	}()

	wg.Wait()
	log.I(TAG, "Server stopped")
}
