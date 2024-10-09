package main

import (
    "flag"
    "github.com/kmiit/vivi/utils"
)

func main () {
    var port string
    
    flag.StringVar(&port, "p", "8080", "端口，默认8080")
    flag.Parse()
    
    utils.RunServer(port)
}
