package utils

import (
    "fmt"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)

func RunServer (p string) {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    port := ":" + p
    fmt.Println("Vivi is listening on: ", p)
    s := &http.Server{
    	Addr:           port,
    	Handler:        router,
    	ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
    	MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
