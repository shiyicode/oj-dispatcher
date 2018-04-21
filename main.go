package main

import (
	"flag"
	"fmt"
	"net/http"

	"runtime"

	"github.com/TV4/graceful"
	"github.com/open-fightcoder/oj-dispatcher/common"
	"github.com/open-fightcoder/oj-dispatcher/common/g"
	"github.com/open-fightcoder/oj-dispatcher/consumer"
	"github.com/open-fightcoder/oj-dispatcher/dispatcher"
	"github.com/open-fightcoder/oj-dispatcher/router"
)

func main() {
	// 获取命令行参数
	cfgFile := flag.String("c", "cfg/cfg.toml.debug", "set config file")
	flag.Parse()

	// 初始化
	common.Init(*cfgFile)

	// 启动路由
	router := router.GetRouter()

	// 启动调度者
	dispatcher.Start(runtime.NumCPU(), runtime.NumCPU())

	// 启动消费者
	comsumer.Start()

	// 优雅退出
	graceful.LogListenAndServe(&http.Server{
		Addr:    fmt.Sprintf(":%d", g.Conf().Run.HTTPPort),
		Handler: router,
	})

	comsumer.Stop()
	dispatcher.Stop()
	common.Close()
}
