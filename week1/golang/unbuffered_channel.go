package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	workers = 2
)

// sử dụng để đợi các goroutine kết thúc
var wg sync.WaitGroup

// có sản phẩm, 2 nhân viên lần lượt lấy sản phẩm
func work(name string, store chan int) {
	defer wg.Done()

	for {
		// đợi sản phẩm
		product, err := <-store
		if !err {
			fmt.Println("Kho het san pham")
			return
		}

		// sản phẩm hết
		if product == int(rand.Int63n(50)) {
			fmt.Printf("%s lay san pham cuoi cung\n", name)

			//đóng kênh
			close(store)
			return
		}

		fmt.Printf("%s lay san pham thu %d\n", name, product)
		product++
		// thêm sản phẩm vào
		store <- product
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	store := make(chan int)
	wg.Add(workers)

	// tạo ra 2 nhân viên
	go work("Nhan vien A", store)
	go work("Nhan vien B", store)

	// bắt đầu cho sản phầm vào
	store <- 1

	wg.Wait()
	fmt.Println("Hoan thanh cong viec")
}
