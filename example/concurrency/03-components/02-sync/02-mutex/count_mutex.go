package main

import (
	"fmt"
	"sync"
)

func Mutex_example() {
	// 定义一个整数变量 count 来表示计数器的值
	var count int
	// 定义一个互斥锁 lock 来保护对 count 变量的访问
	var lock sync.Mutex
	// 定义一个函数 increment 来增加 count 的值
	increment := func() {
		lock.Lock()         // 请求对临界区的独占，用互斥锁解决
		defer lock.Unlock() // 完成对临界区锁定的保护
		count++
		fmt.Printf("Increment: %d\n", count)
	}
	// 定义一个函数 decrement 来减少 count 的值
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrement: %d\n", count)
	}
	// 定义一个 WaitGroup 来等待所有的 goroutine 完成
	var arithmetic sync.WaitGroup
	// 增量
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	// 减量
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()
	fmt.Printf("Final count: %d\n", count)
}
