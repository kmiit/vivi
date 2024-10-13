package server

import (
//    "context"
	"fmt"
	"net/http"
	//"strconv"
	"time"

	"github.com/kmiit/vivi/utils/server/handlers"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	initRouter(r)
	fmt.Println("Vivi is listening on: ", 8080)
	s := &http.Server{
		Addr:           ":8080", // `+ strconv.FormatInt(int64(config.Port), 10),`
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func initRouter(r *gin.Engine) {
	r.Any("/ping", handlers.Pong)
}
