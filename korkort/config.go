package korkort

import (
	"github.com/Unknwon/goconfig"
	"log"
	"os"
)

var (
	_CFG_PATH        = "conf/app.ini"
	_CFG_CUSTOM_PATH = "conf/custom.ini"

	Cfg *goconfig.ConfigFile
)

func InitConfig() {
	var err error

	Cfg, err = goconfig.LoadConfigFile(_CFG_PATH)
	if err != nil {
		log.Fatalf("Fail to load config file(%s): %v", _CFG_PATH, err)
	}

	var f os.FileInfo
	f, err = os.Stat(_CFG_CUSTOM_PATH)
	if err == nil && !f.IsDir() {
		if err = Cfg.AppendFiles(_CFG_CUSTOM_PATH); err != nil {
			log.Fatalf("Fail to load config file(%s): %v", _CFG_CUSTOM_PATH, err)
		}
	}
}
