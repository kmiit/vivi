package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kmiit/vivi/utils/server/handlers"
)

const TAG = "Server"

func initRouter(r *gin.Engine) {
	r.Any("/ping", handlers.Pong) // Test backend server is running or not
}
