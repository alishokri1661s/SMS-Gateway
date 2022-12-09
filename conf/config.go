package conf

import (
	"log"

	"github.com/go-ini/ini"
)

type database struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     string
}

type server struct {
	Host   string
	Port string
}

var DatabaseSetting = &database{}
var ServerSetting = &server{}

func Setup() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'app.ini': %v", err)
		return
	}

	err = cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'database' section: %v", err)
		return
	}

	err = cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'server' section: %v", err)
		return
	}
}
