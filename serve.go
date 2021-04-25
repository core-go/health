package health

import (
	"log"
	"net/http"
	"strconv"
)

type ServerConf struct {
	Name         string `mapstructure:"name" json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Version      string `mapstructure:"version" json:"version,omitempty" gorm:"column:version" bson:"version,omitempty" dynamodbav:"version,omitempty" firestore:"version,omitempty"`
	Port         *int64 `mapstructure:"port" json:"port,omitempty" gorm:"column:port" bson:"port,omitempty" dynamodbav:"port,omitempty" firestore:"port,omitempty"`
	Secure       bool   `mapstructure:"secure" json:"secure,omitempty" gorm:"column:secure" bson:"secure,omitempty" dynamodbav:"secure,omitempty" firestore:"secure,omitempty"`
	Log          *bool  `mapstructure:"log" json:"log,omitempty" gorm:"column:log" bson:"log,omitempty" dynamodbav:"log,omitempty" firestore:"log,omitempty"`
	Monitor      *bool  `mapstructure:"monitor" json:"monitor,omitempty" gorm:"column:monitor" bson:"monitor,omitempty" dynamodbav:"monitor,omitempty" firestore:"monitor,omitempty"`
	CORS         *bool  `mapstructure:"cors" json:"cors,omitempty" gorm:"column:cors" bson:"cors,omitempty" dynamodbav:"cors,omitempty" firestore:"cors,omitempty"`
	AllowOrigin  string `mapstructure:"allow_origin" json:"allowOrigin,omitempty" gorm:"column:alloworigin" bson:"allowOrigin,omitempty" dynamodbav:"allowOrigin,omitempty" firestore:"allowOrigin,omitempty"`
	AllowMethods string `mapstructure:"allow_methods" json:"allowMethods,omitempty" gorm:"column:allowMethods" bson:"allowMethods,omitempty" dynamodbav:"allowMethods,omitempty" firestore:"allowMethods,omitempty"`
	AllowHeaders string `mapstructure:"allow_headers" json:"allowHeaders,omitempty" gorm:"column:allowHeaders" bson:"allowHeaders,omitempty" dynamodbav:"allowHeaders,omitempty" firestore:"allowHeaders,omitempty"`
}

func Serve(conf ServerConf, handler *HealthHandler) {
	server := ""
	if conf.Port != nil && *conf.Port > 0 {
		server = ":" + strconv.FormatInt(*conf.Port, 10)
	}
	if len(conf.Version) > 0 {
		if conf.Port != nil && *conf.Port > 0 {
			log.Println("Start https service: " + conf.Name + " at port " + strconv.FormatInt(*conf.Port, 10) + " with version " + conf.Version)
		} else {
			log.Println("Start https service: " + conf.Name + " with version " + conf.Version)
		}
	} else {
		if conf.Port != nil && *conf.Port > 0 {
			log.Println("Start https service: " + conf.Name + " at port " + strconv.FormatInt(*conf.Port, 10))
		} else {
			log.Println("Start https service: " + conf.Name)
		}
	}
	http.HandleFunc("/health", handler.Check)
	http.HandleFunc("/", handler.Check)
	http.ListenAndServe(server, nil)
}