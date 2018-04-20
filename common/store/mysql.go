package store

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/open-fightcoder/oj-dispatcher/common/g"
	log "github.com/sirupsen/logrus"
)

var OrmWeb *xorm.Engine

func InitMysql() {
	var err error

	conf := g.Conf()

	// Web端数据库Orm引擎
	{
		OrmWeb, err = xorm.NewEngine("mysql", conf.Mysql.WebAddr)

		if err != nil {
			log.Fatalln("fail to connect mysql", conf.Mysql.WebAddr, err)
		}

		OrmWeb.SetMaxIdleConns(conf.Mysql.MaxIdle)
		OrmWeb.SetMaxOpenConns(conf.Mysql.MaxOpen)

		if conf.Mysql.Debug {
			OrmWeb.ShowSQL(true)
			OrmWeb.ShowExecTime(true)
			OrmWeb.Logger().SetLevel(core.LOG_DEBUG)
		} else {
			OrmWeb.Logger().SetLevel(core.LOG_WARNING)
		}
	}
}

func CloseMysql() {

}
