package common

import (
	"github.com/open-fightcoder/oj-dispatcher/common/g"
	"github.com/open-fightcoder/oj-dispatcher/common/store"
)

func Init(cfgFile string) {
	g.LoadConfig(cfgFile)
	g.InitLog()
	store.InitMysql()
	store.InitRedis()
}

func Close() {
	store.CloseMysql()
	store.CloseRedis()
}
