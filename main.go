package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/TV4/graceful"
	"github.com/open-fightcoder/oj-dispatcher/common"
	"github.com/open-fightcoder/oj-dispatcher/common/g"
	"github.com/open-fightcoder/oj-dispatcher/consumer"
	"github.com/open-fightcoder/oj-dispatcher/router"
)

func main() {
	cfgFile := flag.String("c", "cfg/cfg.toml.debug", "set config file")
	flag.Parse()

	common.Init(*cfgFile)
	defer common.Close()

	router := router.GetRouter()

	comsumer.Start()

	graceful.LogListenAndServe(&http.Server{
		Addr:    fmt.Sprintf(":%d", g.Conf().Run.HTTPPort),
		Handler: router,
	})
}
