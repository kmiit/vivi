package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmiit/vivi/utils/config"
	"github.com/kmiit/vivi/utils/log"
)

func InitServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	initRouter(r)
	address := config.ServerConfig.Address + ":" + config.ServerConfig.Port
	log.I(TAG, "vivi is listening on: ", address)
	s := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.F(TAG, err)
	}
}
