package postgres

import (
	"fmt"
	"log"
	"sync"

	"github.com/Golang/week3.1/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var once sync.Once
var db *gorm.DB

// GetController return controller
func GetConnectionPostgres() *gorm.DB {
	once.Do(func() {
		connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.GetConfig().HostnamePostgres, config.GetConfig().PortPostgres, config.GetConfig().UsernamePostgres, config.GetConfig().DbNamePostgres, config.GetConfig().PasswordPostgres)
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
