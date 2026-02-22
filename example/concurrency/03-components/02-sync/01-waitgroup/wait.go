package main

import (
	"fmt"
	"sync"
	"time"
)

func example01() {
	var wg sync.WaitGroup
	// 调用 Add 方法，参数为 1，表示有一个 goroutine 开始了
	wg.Add(1)
	go func() {
		// 使用 defer 关键字来确保在 goroutine 结束时调用 Done 方法，表示这个 goroutine 已经完成了
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1 * time.Second)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2 * time.Second)
	}()
	// 调用 Wait 方法，这将阻塞 main goroutine，等待所有的 goroutine 完成
	wg.Wait()
	fmt.Println("All goroutines finished.")
}

func example02() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from goroutine %v\n", id)
	}
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}

func example03() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from goroutine %v\n", id)
	}
	const numGreeters = 5
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < numGreeters; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			go hello(&wg, id)
			fmt.Printf("Hello from goroutine %v\n", id)
			mu.Unlock()
		}(i + 1)
	}
	wg.Wait()
}

func main() {
	example01()
	example02()
	example03()
}
