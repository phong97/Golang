package main

import (
	"fmt"

	entity "github.com/Golang/week3.1/db/entity"
	mysql "github.com/Golang/week3.1/db/mysql"
	"github.com/Golang/week3.1/db/postgres"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Analytic entity.Analytic

func main() {
	fmt.Println("PostgreSQL:")
	postgres_conn := postgres.GetConnectionPostgres()
	result, err := postgres_conn.Raw("with table1 as ( select p.rental_id, p.amount, r.inventory_id, i.film_id, f.title, f.rating from payment p left join rental r on p.rental_id = r.rental_id left join inventory i on i.inventory_id = r.inventory_id left join film f on f.film_id = i.film_id), table2 as ( select table1.title, table1.rating, sum(table1.amount) as sum from table1 group by 1,2 order by 3 desc ) select * from ( select table2.*, row_number() over(order by table2.sum desc) as rank_per_total, format('%s/%s', row_number() over(partition by table2.rating order by table2.sum desc), count (*) over (partition by table2.rating) ) as rank_per_rating from table2 ) as table3 where table3.rank_per_total <= 10 order by table3.rank_per_total;").Rows()

	if err != nil {
		panic(err)
	}

	for result.Next() {
		var analytic Analytic
		if err = postgres_conn.ScanRows(result, &analytic); err != nil {
			fmt.Println(err)
		}
		fmt.Println(analytic)
	}

	fmt.Println("MySQL:")
	mysql_conn := mysql.GetConnectionMySQL()
	result, err = mysql_conn.Raw("with table1 as ( select p.rental_id, p.amount, r.inventory_id, i.film_id, f.title, f.rating from payment p left join rental r on p.rental_id = r.rental_id left join inventory i on i.inventory_id = r.inventory_id left join film f on f.film_id = i.film_id), table2 as ( select table1.title, table1.rating, sum(table1.amount) as sum from table1 group by 1,2 order by 3 desc ) select * from ( select table2.*, row_number() over(order by table2.sum desc) as rank_per_total, concat( row_number() over(partition by table2.rating order by table2.sum desc), \" / \", count(*) over (partition by table2.rating) ) as rank_per_rating from table2 ) as table3 where table3.rank_per_total <= 10 order by table3.rank_per_total;").Rows()
	defer result.Close()
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var analytic Analytic
		if err = mysql_conn.ScanRows(result, &analytic); err != nil {
			fmt.Println(err)
		}
		fmt.Println(analytic)
	}
}
