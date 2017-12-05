package korkort

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	_CFG_PATH = "../conf/app.ini"
	_CFG_CUSTOM_PATH = ""

	InitConfig()
	dbFile, err := Cfg.GetValue("app", "database")

	if err != nil {
		t.Error("Expected app.ini be loaded")
	}

	if dbFile != "data/main.db" {
		t.Error("Expected app.database == 'data/main.db'")
	}

}
