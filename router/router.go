package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-dispatcher/common/g"
	"github.com/open-fightcoder/oj-dispatcher/router/controllers/apiv1"
	"github.com/open-fightcoder/oj-dispatcher/router/controllers/authv1"
	"github.com/open-fightcoder/oj-dispatcher/router/middleware"
)

var router *gin.Engine
var once sync.Once

// 获取路由并初始化
func GetRouter() *gin.Engine {
	// 只执行一次
	once.Do(func() {
		initRouter()
	})
	return router
}

// 初始化路由
func initRouter() {
	router = gin.Default()

	router.Use(middleware.Cors())
	router.Use(middleware.Recovery())
	router.Use(middleware.MaxAllowed(g.Conf().Run.MaxAllowed))

	authRouter := router.Group("/authv1", middleware.Auth())
	authv1.Register(authRouter)

	apiRouter := router.Group("/apiv1")
	apiv1.Register(apiRouter)
}
