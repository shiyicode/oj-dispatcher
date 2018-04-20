package g

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Run   RunConfig   `toml:"run"`
	Log   LogConfig   `toml:"log"`
	Mysql MysqlConfig `toml:"mysql"`
	Jwt   JwtConfig   `toml:"jwt"`
	Nsq   NsqConfig   `toml:"nsq"`
}

type RunConfig struct {
	WaitTimeout int    `toml:"waitTimeout"`
	HTTPPort    int    `toml:"httpPort"`
	Mode        string `toml:"mode"`
	MaxAllowed  int    `toml:"maxAllowed"`
}

type LogConfig struct {
	Enable    bool   `toml:"enable"`
	Path      string `toml:"path"`
	Level     string `toml:"level"`
	RotatTime int    `toml:"rotatTime"`
	MaxAge    int    `toml:"maxAge"`
}

type MysqlConfig struct {
	MaxIdle int    `toml:"maxIdle"`
	MaxOpen int    `toml:"maxOpen"`
	Debug   bool   `toml:"debug"`
	WebAddr string `toml:"webAddr"`
}

type JwtConfig struct {
	EncodeMethod     string `toml:"encodeMethod"`
	MaxEffectiveTime int64  `toml:"maxEffectiveTime"`
}

type NsqConfig struct {
	Lookupds     []string `toml:"lookupds"`
	JudgeTopic   string   `toml:"judgeTopic"`
	JudgeChannel string   `toml:"judgeChannel"`
	HandlerCount int      `toml:"handlerCount"`
}

var (
	ConfigFile string
	config     *Config
	configLock = new(sync.RWMutex)
)

func Conf() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

// 加载配置文件
func LoadConfig(cfgFile string) {
	configLock.Lock()
	defer configLock.Unlock()

	// 配置文件路径是否为空
	if cfgFile == "" {
		log.Fatalln("config file not specified: use -c $filename")
	}

	// 配置文件是否存在
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		log.Fatalln("config file specified not found:", cfgFile)
	}

	ConfigFile = cfgFile

	if bs, err := ioutil.ReadFile(cfgFile); err != nil {
		log.Fatalf("read config file failed: %s\n", err.Error())
	} else {
		if _, err := toml.Decode(string(bs), &config); err != nil {
			log.Fatalf("decode config file failed: %s\n", err.Error())
		} else {
			log.Printf("load config from %s\n", cfgFile)
			log.Printf("config: %#v\n", config)
		}
	}
}
