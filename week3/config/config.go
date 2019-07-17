package config

import (
	"sync"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	UsernameDb string
	PasswordDb string
	HostnameDb string
	PortDb     string
	DbName     string
}

var once sync.Once
var config *ServerConfig

func GetConfig() *ServerConfig {
	once.Do(func() {
		config = readConfig()
	})

	return config
}

func readConfig() *ServerConfig {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetDefault("database.dbname", "test")
	viper.SetDefault("database.username", "postgres")
	viper.SetDefault("database.password", "12345")
	viper.SetDefault("database.hostname", "localhost")
	viper.SetDefault("database.port", "5432")

	conf := ServerConfig{}
	//db info
	conf.DbName = viper.GetString("database.dbname")
	conf.UsernameDb = viper.GetString("database.username")
	conf.PasswordDb = viper.GetString("database.password")
	conf.HostnameDb = viper.GetString("database.hostname")
	conf.PortDb = viper.GetString("database.port")
	return &conf
}
