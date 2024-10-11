package main

import (
    "context"
    "fmt"
    "sync"

	"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils"
)

func run(config types.ServerConfig) {
    var wg sync.WaitGroup
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
	wg.Add(1)
	go func() {
	    defer wg.Done()
	    utils.RunServer(config, ctx)
	}()
	
	wg.Wait()
	fmt.Println("Done")
}
