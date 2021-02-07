package health

import (
	"log"
	"net/http"
	"strconv"
)

type ServerConf struct {
	Name       string `mapstructure:"name"`
	Version    string `mapstructure:"version"`
	Port       int64  `mapstructure:"port"`
	CORS       bool   `mapstructure:"cors"`
	Logging    bool   `mapstructure:"logging"`
	Monitoring bool   `mapstructure:"monitoring"`
	Secure     bool   `mapstructure:"secure"`
}

func Serve(conf ServerConf, handler *HealthHandler) {
	server := ""
	if conf.Port > 0 {
		server = ":" + strconv.FormatInt(conf.Port, 10)
	}
	if len(conf.Version) > 0 {
		if conf.Port > 0 {
			log.Println("Start https service: " + conf.Name + " at port " + strconv.FormatInt(conf.Port, 10) + " with version " + conf.Version)
		} else {
			log.Println("Start https service: " + conf.Name + " with version " + conf.Version)
		}
	} else {
		if conf.Port > 0 {
			log.Println("Start https service: " + conf.Name + " at port " + strconv.FormatInt(conf.Port, 10))
		} else {
			log.Println("Start https service: " + conf.Name)
		}
	}
	http.HandleFunc("/health", handler.Check)
	http.HandleFunc("/", handler.Check)
	http.ListenAndServe(server, nil)
}
