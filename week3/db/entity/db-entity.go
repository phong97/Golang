package entity

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type Customer struct {
	CustomerID int32          `gorm:"AUTO_INCREMENT; primary_key; unique_index"`
	Firstname  string         `gorm:"type:varchar(30);NOT NULL"`
	Lastname   string         `gorm:"type:varchar(30);NOT NULL"`
	Email      string         `gorm:"type:varchar(255);unique_index"`
	Address    pq.StringArray `gorm:"type:varchar(255)[]"`
	Secrets    postgres.Hstore
	Body       postgres.Jsonb
}
