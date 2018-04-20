package middleware

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-dispatcher/common/components"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		data, err := base64.URLEncoding.DecodeString(token)
		if err != nil {
			panic(err)
		}

		token = string(data)

		if token != "" {
			if flag, userId := components.RequireTokenAuthentication(token); flag == true {
				c.Set("userId", userId)
				c.Next()
			}
		}
		c.JSON(http.StatusForbidden, "authv1 failure")
	}
}
