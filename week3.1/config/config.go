package config

import (
	"sync"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	UsernameMySQL    string
	PasswordMySQL    string
	HostnameMySQL    string
	PortMySQL        string
	DbNameMySQL      string
	UsernamePostgres string
	PasswordPostgres string
	HostnamePostgres string
	PortPostgres     string
	DbNamePostgres   string
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
	//set default
	viper.SetDefault("postgres.dbname", "dvdrental")
	viper.SetDefault("postgres.username", "postgres")
	viper.SetDefault("postgres.password", "12345")
	viper.SetDefault("postgres.hostname", "localhost")
	viper.SetDefault("postgres.port", "5432")

	viper.SetDefault("mysql.dbname", "sakila")
	viper.SetDefault("mysql.username", "root")
	viper.SetDefault("mysql.password", "12345")
	viper.SetDefault("mysql.hostname", "localhost")
	viper.SetDefault("mysql.port", "3306")

	conf := ServerConfig{}
	//postgres info
	conf.DbNamePostgres = viper.GetString("postgres.dbname")
	conf.UsernamePostgres = viper.GetString("postgres.username")
	conf.PasswordPostgres = viper.GetString("postgres.password")
	conf.HostnamePostgres = viper.GetString("postgres.hostname")
	conf.PortPostgres = viper.GetString("postgres.port")
	//mysql info
	conf.DbNameMySQL = viper.GetString("mysql.dbname")
	conf.UsernameMySQL = viper.GetString("mysql.username")
	conf.PasswordMySQL = viper.GetString("mysql.password")
	conf.HostnameMySQL = viper.GetString("mysql.hostname")
	conf.PortMySQL = viper.GetString("mysql.port")
	return &conf
}
