package base

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserId(c *gin.Context) int64 {
	str, exists := c.Get("userId")
	if exists == false {
		return 0
	}
	id, _ := strconv.ParseInt(str.(string), 10, 64)
	return id
}
