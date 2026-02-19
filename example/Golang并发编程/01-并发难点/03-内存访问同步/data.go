package main

import (
	"fmt"
	"sync"
)

func example01() {
	// 初始化 data 变量
	var data int

	// 启动一个新的 goroutine 来修改 data 的值
	go func() {
		data++
	}()
	// 主 goroutine 继续执行，可能会在 data 被修改之前检查它的值
	if data == 0 {
		fmt.Println("the value is 0")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}

func example02() {
	// 初始化 memoryAccess 互斥锁和 value 变量
	var memoryAccess sync.Mutex
	var value int

	// 启动一个新的 goroutine 来修改 value 的值
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	// 加锁以确保在检查 value 的值时不会被其他 goroutine 修改
	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}
