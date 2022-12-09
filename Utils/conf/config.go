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
	Host string
	Port string
}

type messageBroker struct {
	User     string
	Password string
	Host     string
	Port     string
}

var DatabaseSetting = &database{}
var ServerSetting = &server{}
var MessageBrokerSetting = &messageBroker{}
var GhasedakAPIKey string

func Setup() {
	cfg, err := ini.Load("Utils/conf/app.ini")
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

	err = cfg.Section("messageBroker").MapTo(MessageBrokerSetting)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'messageBroker' section: %v", err)
		return
	}

	key, err := cfg.Section("ghasedak").GetKey("api_key")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'ghasedak' section: %v", err)
		return
	}
	GhasedakAPIKey = key.Value()
}
