package config

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var C *Config

type Config struct {
	Basic         Basic         `mapstructure:"basic"`
	Logger        Logger        `mapstructure:"logger"`
	HTTPServer    HTTPServer    `mapstructure:"http_server"`
	Database      SQLDatabase   `mapstructure:"database"`
	TraceProvider TraceProvider `mapstructure:"trace_provider"`
}

type Basic struct {
	CORSWhiteList            []string   `mapstructure:"cors_white_list"`
	PopularPostsFromLastDays int        `mapstructure:"popular_posts_from_last_days"`
	Pagination               Pagination `mapstructure:"pagination"`
}

type Pagination struct {
	MaximumBlogPerPage int `mapstructure:"maximum_blog_per_page"`
}

type Logger struct {
	Level string `mapstructure:"level"`
}

func (c *Config) String() string {
	return fmt.Sprintf("HTTP\t\t|%s\n", c.HTTPServer)
}

type HTTPServer struct {
	Listen            string        `mapstructure:"listen"`
	ReadTimeout       time.Duration `mapstructure:"read_Timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
	IdleTimeout       time.Duration `mapstructure:"idle_timeout"`
}

func (i HTTPServer) String() string {
	return fmt.Sprintf("Listen to: %s", i.Listen)
}

type SQLDatabase struct {
	Driver        string        `mapstructure:"driver"`
	Host          string        `mapstructure:"host"`
	Port          int           `mapstructure:"port"`
	DB            string        `mapstructure:"db"`
	User          string        `mapstructure:"user"`
	Password      string        `mapstructure:"password"`
	TimeZone      string        `mapstructure:"time_zone"`
	MaxConn       int           `mapstructure:"max_conn"`
	IdleConn      int           `mapstructure:"idle_conn"`
	Timeout       time.Duration `mapstructure:"timeout"`
	DialRetry     int           `mapstructure:"dial_retry"`
	DialTimeout   time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout   time.Duration `mapstructure:"read_timeout"`
	WriteTimeout  time.Duration `mapstructure:"write_timeout"`
	UpdateTimeout time.Duration `mapstructure:"update_timeout"`
	DeleteTimeout time.Duration `mapstructure:"delete_timeout"`
	QueryTimeout  time.Duration `mapstructure:"query_timeout"`
}

func (d SQLDatabase) DSN() (dsn string) {
	switch d.Driver {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", d.Host, d.User, d.Password, d.DB, d.Port, d.TimeZone)
	default:
		log.Fatalf("SQLDatabase driver is not supported: %s", d.Driver)
	}

	return
}

// String represents most usable values for the SQLDatabase config
func (d SQLDatabase) String() (str string) {
	switch d.Driver {
	case "postgres":
		str = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", d.Host, d.User, d.Password, d.DB, d.Port, d.TimeZone)
	default:
		log.Fatalf("SQLDatabase driver is not supported: %s", d.Driver)
	}

	return
}

type TraceProvider struct {
	ServiceName           string  `mapstructure:"service_name"`
	DeploymentEnvironment string  `mapstructure:"deployment_environment"`
	AgentHost             string  `mapstructure:"agent_host"`
	AgentPort             string  `mapstructure:"agent_port"`
	SamplerRatio          float64 `mapstructure:"sampler_ratio"`
}

func Init(filename string) {
	c := new(Config)
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigFile(filename)
	v.SetConfigType("yml")

	err := v.ReadInConfig()
	if err != nil {
		log.WithError(err).Fatalf("failed on config `%s` unmarshal", filename)
	}

	err = v.Unmarshal(c)
	if err != nil {
		log.WithError(err).Fatalf("load configurations failed")
	}

	C = c
	log.Infof("config file opened successfully. filename: %s", filename)
}
