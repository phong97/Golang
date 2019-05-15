package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	workers  = 4
	products = 10
)

// sử dụng để đợi các goroutine kết thúc
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

		// random thời gian cho việc hoàn thành sản phầm
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("%s hoan thanh san phan %d\n", name, product)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	store := make(chan int)
	wg.Add(workers)

	// tạo 4 nhân viên
	for worker := 1; worker <= workers; worker++ {
		go work(fmt.Sprintf("Nhan vien %d", worker), store)
	}

	// thêm nhiều sản phẩm vào kênh
	for product := 1; product <= products; product++ {
		store <- product
	}
	close(store)

	wg.Wait()
	fmt.Println("Tat ca nhan vien hoan thanh cong viec")
}
