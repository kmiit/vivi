package server

import (
//    "context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils/server/handlers"

	"github.com/gin-gonic/gin"
)

func RunServer(config types.Config) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	initRouter(r)
	address := config.Server.Address + ":" + strconv.FormatInt(int64(config.Server.Port), 10)
	fmt.Println("Vivi is listening on: ", address)
	s := &http.Server{
		Addr:           address,
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
