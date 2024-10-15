package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils/log"
)

func RunServer(config types.Config) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	initRouter(r)
	address := config.Server.Address + ":" + strconv.FormatInt(int64(config.Server.Port), 10)
	log.I(TAG, "vivi is listening on: ", address)
	s := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}