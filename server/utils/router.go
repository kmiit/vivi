package utils

import (
    "context"
	"fmt"
	"net/http"
	//"strconv"
	"time"

	//"github.com/kmiit/vivi/types"

	"github.com/gin-gonic/gin"
)

func RunServer(ctx context.Context) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("Vivi is listening on: ", 8080)
	s := &http.Server{
		Addr:           ":8080", // `+ strconv.FormatInt(int64(config.Port), 10),`
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
