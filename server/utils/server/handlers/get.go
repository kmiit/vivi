package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kmiit/vivi/utils/db"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

func Get(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		items, err := db.GetAllPublic(ctx, db.FILE_NAMESPACE)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, items)
		return
	} else {
		res, err := db.GetPublic(ctx, db.FILE_NAMESPACE + id)
		if err.Error() == "key does not exist" {
			c.JSON(404, gin.H{"error": "Invalid ID"})
		} else if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, res)
		}
		return
	}
}
