package main

import (
	"encoding/json"
	"fmt"
	"time"

	client "github.com/Golang/week3/db"
	entity "github.com/Golang/week3/db/entity"

	"github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const tbCustomer = "customer"

type Customer entity.Customer

func (customer Customer) Create() error {
	db := client.GetConnectionDB()
	return db.Debug().Table(tbCustomer).Create(&customer).Error
}

func (customer *Customer) Read() error {
	db := client.GetConnectionDB()
	return db.Debug().Table(tbCustomer).First(&customer).Error
}

func Update() error {
	db := client.GetConnectionDB()
	return db.Debug().Table(tbCustomer).Where("firstname = 'Bao'").Updates(Customer{Firstname: "ABao"}).Error
}

func Delete() error {
	db := client.GetConnectionDB()
	return db.Debug().Table(tbCustomer).Where("firstname = 'ABao'").Delete(Customer{}).Error
}

func main() {
	db := client.GetConnectionDB()
	db.Debug().AutoMigrate(&entity.Customer{})
	body := json.RawMessage(`{"height": 160, "weight": 50}`)
	cardID := "123456789"

	new_customer := &Customer{
		Firstname: "anh Bao",
		Lastname:  "Chi",
		Address:   []string{"Quan 10", "Quan 11"},
		Email:     fmt.Sprintf("email%v@vng.com.vn", time.Now().Unix()),
		Body:      postgres.Jsonb{body},
		Secrets:   postgres.Hstore{"cardID": &cardID},
	}

	new_customer.Create()
	var customer Customer
	customers.Read()
	for _, c := range customers {
		fmt.Println(c)
	}
	// db.Debug().Table(tbCustomer).Create(&new_customer)

	// var customer entity.Customer
	// // var customers []entity.Customer

	// db.Debug().Table(tbCustomer).First(&customer)
	// if err := db.Debug().Table(tbCustomer).First(&customer); err != nil {
	// 	fmt.Println(customer)
	// }

	// var _note map[string]interface{}
	// b, _ := customer.Note.MarshalJSON()
	// json.Unmarshal(b, &_note)
	// fmt.Println(_note)
}
