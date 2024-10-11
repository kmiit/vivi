package cmd

import (
    "context"
    "fmt"
    "sync"

	//"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the vivi server",
	Long:  `Run the vivi server in frontend..`,
  	Run: func(cmd *cobra.Command, args []string) {
    	runServer()
  },
}

func runServer() {
    var wg sync.WaitGroup
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
	wg.Add(1)
	go func() {
	    defer wg.Done()
	    utils.RunServer(ctx)
	}()
	
	wg.Wait()
	fmt.Println("Done")
}
