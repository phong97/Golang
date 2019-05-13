package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	workers = 4
	products = 10
)

var wg sync.WaitGroup

func work(name string, store chan int) {
	defer wg.Done()

	for {
		product, err := <-store
		if !err {
			fmt.Printf("%s hoan thanh cong viec\n", name)
			return
		}

		fmt.Printf("%s lay san pham thu %d\n", name, product)
		time.Sleep(1000*time.Millisecond)
		fmt.Printf("%s hoan thanh san phan %d\n", name, product)
	}
}

func main() {
	store := make(chan int)
	wg.Add(workers)

	for worker := 1; worker <= workers; worker++ {
		go work(fmt.Sprintf("Nhan vien %d", worker), store)
	}

	for product := 1; product <= products; product++ {
		store <- product
	}
	close(store)

	wg.Wait()
	fmt.Println("Tat ca nhan vien hoan thanh cong viec")
}