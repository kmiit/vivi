package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kmiit/vivi/utils/db"
	"github.com/kmiit/vivi/utils/storage"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

func Get(c *gin.Context) {
	id := c.Query("id")
	items, err := db.GetAllOuter(ctx, storage.FILE_NAMESPACE)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		if id == "" {
			c.JSON(200, items)
			return
		}
		for _, item := range items {
			idInt, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid ID"})
				return
			}
			if item.ID == idInt {
				c.JSON(200, item)
				break
			}
		}
	}
}
