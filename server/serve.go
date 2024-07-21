package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ServerConfig struct {
	AppId             string         `yaml:"app_id" mapstructure:"app_id" json:"appId,omitempty" gorm:"column:appid" bson:"appId,omitempty" dynamodbav:"appId,omitempty" firestore:"appId,omitempty"`
	Name              string         `yaml:"name" mapstructure:"name" json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Version           string         `yaml:"version" mapstructure:"version" json:"version,omitempty" gorm:"column:version" bson:"version,omitempty" dynamodbav:"version,omitempty" firestore:"version,omitempty"`
	Port              *int64         `yaml:"port" mapstructure:"port" json:"port,omitempty" gorm:"column:port" bson:"port,omitempty" dynamodbav:"port,omitempty" firestore:"port,omitempty"`
	Secure            bool           `yaml:"secure" mapstructure:"secure" json:"secure,omitempty" gorm:"column:secure" bson:"secure,omitempty" dynamodbav:"secure,omitempty" firestore:"secure,omitempty"`
	Log               *bool          `yaml:"log" mapstructure:"log" json:"log,omitempty" gorm:"column:log" bson:"log,omitempty" dynamodbav:"log,omitempty" firestore:"log,omitempty"`
	Monitor           *bool          `yaml:"monitor" mapstructure:"monitor" json:"monitor,omitempty" gorm:"column:monitor" bson:"monitor,omitempty" dynamodbav:"monitor,omitempty" firestore:"monitor,omitempty"`
	CORS              *bool          `yaml:"cors" mapstructure:"cors" json:"cors,omitempty" gorm:"column:cors" bson:"cors,omitempty" dynamodbav:"cors,omitempty" firestore:"cors,omitempty"`
	WriteTimeout      *time.Duration `yaml:"write_timeout" mapstructure:"write_timeout" json:"writeTimeout,omitempty" gorm:"column:writetimeout" bson:"writeTimeout,omitempty" dynamodbav:"writeTimeout,omitempty" firestore:"writeTimeout,omitempty"`
	ReadTimeout       *time.Duration `yaml:"read_timeout" mapstructure:"read_timeout" json:"readTimeout,omitempty" gorm:"column:readtimeout" bson:"readTimeout,omitempty" dynamodbav:"readTimeout,omitempty" firestore:"readTimeout,omitempty"`
	ReadHeaderTimeout *time.Duration `yaml:"read_header_timeout" mapstructure:"read_header_timeout" json:"readHeaderTimeout,omitempty" gorm:"column:readheadertimeout" bson:"readHeaderTimeout,omitempty" dynamodbav:"readHeaderTimeout,omitempty" firestore:"readHeaderTimeout,omitempty"`
	IdleTimeout       *time.Duration `yaml:"idle_timeout" mapstructure:"idle_timeout" json:"idleTimeout,omitempty" gorm:"column:idletimeout" bson:"idleTimeout,omitempty" dynamodbav:"idleTimeout,omitempty" firestore:"idleTimeout,omitempty"`
	MaxHeaderBytes    *int           `yaml:"max_header_bytes" mapstructure:"max_header_bytes" json:"maxHeaderBytes,omitempty" gorm:"column:maxheaderbytes" bson:"maxHeaderBytes,omitempty" dynamodbav:"maxHeaderBytes,omitempty" firestore:"maxHeaderBytes,omitempty"`
	Cert              string         `yaml:"cert" mapstructure:"cert" json:"cert,omitempty" gorm:"column:cert" bson:"cert,omitempty" dynamodbav:"cert,omitempty" firestore:"cert,omitempty"`
	Key               string         `yaml:"key" mapstructure:"key" json:"key,omitempty" gorm:"column:key" bson:"key,omitempty" dynamodbav:"key,omitempty" firestore:"key,omitempty"`
}

func Addr(port *int64) string {
	server := ""
	if port != nil && *port >= 0 {
		server = ":" + strconv.FormatInt(*port, 10)
	}
	return server
}
func ServerInfo(conf ServerConfig) string {
	if len(conf.Version) > 0 {
		if conf.Port != nil && *conf.Port >= 0 {
			return "Start service: " + conf.Name + " at port " + strconv.FormatInt(*conf.Port, 10) + " with version " + conf.Version
		} else {
			return "Start service: " + conf.Name + " with version " + conf.Version
		}
	} else {
		if conf.Port != nil && *conf.Port >= 0 {
			return "Start service: " + conf.Name + " at port " + strconv.FormatInt(*conf.Port, 10)
		} else {
			return "Start service: " + conf.Name
		}
	}
}
func Serve(conf ServerConfig, check func(w http.ResponseWriter, r *http.Request), options ...*tls.Config) {
	log.Println(ServerInfo(conf))
	http.HandleFunc("/health", check)
	http.HandleFunc("/", check)
	srv := CreateServer(conf, nil, options...)
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
}
func CreateServer(conf ServerConfig, handler http.Handler, options ...*tls.Config) *http.Server {
	addr := Addr(conf.Port)
	srv := http.Server{
		Addr:      addr,
		Handler:   nil,
		TLSConfig: nil,
	}
	if len(options) > 0 && options[0] != nil {
		srv.TLSConfig = options[0]
	}
	if conf.ReadTimeout != nil {
		srv.ReadTimeout = *conf.ReadTimeout
	}
	if conf.ReadHeaderTimeout != nil {
		srv.ReadHeaderTimeout = *conf.ReadHeaderTimeout
	}
	if conf.WriteTimeout != nil {
		srv.WriteTimeout = *conf.WriteTimeout
	}
	if conf.IdleTimeout != nil {
		srv.IdleTimeout = *conf.IdleTimeout
	}
	if conf.MaxHeaderBytes != nil && *conf.MaxHeaderBytes > 0 {
		srv.MaxHeaderBytes = *conf.MaxHeaderBytes
	}
	srv.Handler = handler
	return &srv
}
