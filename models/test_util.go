/**
 * Created by leeezm on 2018/4/18.
 * Email: shiyi@fightcoder.com
 */

package models

import (
	"github.com/open-fightcoder/oj-dispatcher/common"
	"sync"
)

var once sync.Once

func InitAllInTest() {
	once.Do(func() {
		common.Init("../cfg/cfg.toml.debug")
	})
}
