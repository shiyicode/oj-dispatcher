package authv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-dispatcher/router/controllers/base"
)

func RegisterAccount(router *gin.RouterGroup) {
	router.POST("quit", httpHandlerQuit)
}

func httpHandlerQuit(c *gin.Context) {
	cookie := http.Cookie{Name: "token", Path: "/", MaxAge: -1}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, base.Success())
}
