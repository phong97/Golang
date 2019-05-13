package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	workers = 2
	products = 10
)

var wg sync.WaitGroup

func work(name string, store chan int) {
	defer wg.Done()

	for {
		product, err := <-store
		if !err {
			fmt.Println("Kho het san pham")
			return
		}

		if product == products {
			fmt.Printf("%s lay san pham cuoi cung\n", name)
			close(store)
			return
		}

		fmt.Printf("%s lay san pham thu %d\n", name, product)
		product++
		store <- product
		time.Sleep(1000 * time.Millisecond)
	}
}


func main() {
	store := make(chan int)
	wg.Add(workers)

	go work("Nhan vien A", store)
	go work("Nhan vien B", store)

	store <- 1

	wg.Wait()
	fmt.Println("Hoan thanh cong viec")
}
