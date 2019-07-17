package main

import (
	"encoding/json"
	"fmt"
	"time"

	client "github.com/Golang/week3/db"
	"github.com/Golang/week3/db/entity"

	"github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const tbCustomer = "customer"

type Customer entity.Customer

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

	fmt.Println("Create: ")
	err := db.Table(tbCustomer).Create(&new_customer).Error
	if err != nil {
		fmt.Println("That bai")
	} else {
		fmt.Println("Thanh cong")
	}

	fmt.Println("Update: ")
	err = db.Table(tbCustomer).Where("firstname = 'anh Bao'").Updates(Customer{Firstname: "ABao"}).Error
	if err != nil {
		fmt.Println("That bai")
	} else {
		fmt.Println("Thanh cong")
	}

	fmt.Println("Read: ")
	var customer Customer
	err = db.Table(tbCustomer).First(&customer).Error
	if err != nil {
		fmt.Println("That bai")
	} else {
		fmt.Println("First name: " + customer.Firstname)
		fmt.Println("Last name: " + customer.Lastname)
		fmt.Println("Address: ")
		for _, value := range customer.Address {
			fmt.Println(value)
		}
		fmt.Println("Body: ")
		var _body map[string]interface{}
		byte_body, _ := customer.Body.MarshalJSON()
		json.Unmarshal(byte_body, &_body)
		fmt.Println(_body)
		fmt.Println("Secrets: ")
		for key, val := range customer.Secrets {
			if val != nil {
				fmt.Println(key + " => " + *val)
			}
		}
	}

	fmt.Println("Delete: ")
	err = db.Table(tbCustomer).Where("firstname = 'ABao'").Delete(Customer{}).Error
	if err != nil {
		fmt.Println("That bai")
	} else {
		fmt.Println("Thanh cong")
	}
}
