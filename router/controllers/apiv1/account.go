package apiv1

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-dispatcher/managers"
	"github.com/open-fightcoder/oj-dispatcher/router/controllers/base"
)

func RegisterAccount(router *gin.RouterGroup) {
	router.POST("account/login", httpHandlerLogin)
	router.POST("account/register", httpHandlerRegister)
}

type AccountParam struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func httpHandlerLogin(c *gin.Context) {
	account := AccountParam{}
	err := c.Bind(&account)
	if err != nil {
		panic(err)
	}
	if flag, token, mess := managers.AccountLogin(account.Email, account.Password); flag == false {
		c.JSON(http.StatusOK, base.Fail(mess))
	} else {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    base64.StdEncoding.EncodeToString([]byte(token)),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, base.Success())
	}
}

func httpHandlerRegister(c *gin.Context) {
	account := AccountParam{}
	err := c.Bind(&account)
	if err != nil {
		panic(err)
	}
	if flag, userId, mess := managers.AccountRegister(account.Email, account.Password); flag == false {
		c.JSON(http.StatusOK, base.Fail(mess))
	} else {
		c.JSON(http.StatusOK, base.Success(userId))
	}
}
