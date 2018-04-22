package base

import "github.com/gin-gonic/gin"

func UserId(c *gin.Context) int64 {
	id, isExit := c.Get("userId")
	if isExit == true {
		num := id.(int64)
		return num
	}
	return 0
}
