package mysql

import (
	"fmt"
	"log"
	"sync"

	"github.com/Golang/week3.1/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var once sync.Once
var db *gorm.DB

// GetController return controller
func GetConnectionMySQL() *gorm.DB {
	once.Do(func() {
		connString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.GetConfig().UsernameMySQL, config.GetConfig().PasswordMySQL, config.GetConfig().HostnameMySQL, config.GetConfig().PortMySQL, config.GetConfig().DbNameMySQL)
		gormDb, err := gorm.Open("mysql", connString)
		// Use singular table as default
		gormDb.SingularTable(true)
		// Keep connection alive
		gormDb.DB().SetConnMaxLifetime(-1)
		// Unlimited open connections
		gormDb.DB().SetMaxOpenConns(-1)
		db = gormDb
		if err != nil {
			log.Fatal(err)
		}
	})
	return db
}
