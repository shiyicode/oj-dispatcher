package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-dispatcher/common/g"
)

func Logger() gin.HandlerFunc {
	conf := g.Conf()
	if conf.Log.Enable {
		logPath := conf.Log.Path
		maxAge := time.Duration(conf.Log.MaxAge)
		rotationTime := time.Duration(conf.Log.RotatTime)
		writer := g.GetLogWriter(logPath, "access", maxAge, rotationTime)
		return gin.LoggerWithWriter(writer)
	}
	return gin.Logger()
}
