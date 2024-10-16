package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kmiit/vivi/utils/server/handlers"
)

const TAG = "Server"

/*
 *\/ping： Check if the server is running.
 *\/get: /get?id=  Get files by `id`, if id is empty, return all files.
 */

func initRouter(r *gin.Engine) {
	r.Any("/ping", handlers.Pong)
	r.GET("/get", handlers.Get)
}
