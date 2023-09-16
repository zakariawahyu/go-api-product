package config

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"time"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Api product config path")
}

type Config struct {
	AppVersion string
	Server     Server
	Mysql      Mysql
	Logger     Logger
}

type Server struct {
	Port         string
	Development  bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Mysql struct {
	Host              string
	Port              string
	User              string
	Password          string
	DbName            string
	MaxIdleConnection int
	MaxOpenConnection int
}

type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv("CONFIG_PATH")
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost != "" {
		cfg.Mysql.Host = mysqlHost
	}

	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort != "" {
		cfg.Mysql.Port = mysqlPort
	}

	return cfg, nil
}
