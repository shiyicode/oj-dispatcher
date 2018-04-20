package apiv1

import "github.com/gin-gonic/gin"

func Register(router *gin.RouterGroup) {
	RegisterSelf(router)
	RegisterAccount(router)
}
