package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/Golang/week3/config"
	"github.com/jinzhu/gorm"
)

var once sync.Once
var db *gorm.DB

// GetController return controller
func GetConnectionDB() *gorm.DB {
	once.Do(func() {
		connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.GetConfig().HostnameDb, config.GetConfig().PortDb, config.GetConfig().UsernameDb, config.GetConfig().DbName, config.GetConfig().PasswordDb)
		gormDb, err := gorm.Open("postgres", connString)
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
