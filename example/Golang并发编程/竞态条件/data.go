package main

import (
	"fmt"
	"time"
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
		fmt.Printf("the value is %v.\n", data)
	}
}

func example02() {
	// 初始化 data 变量
	var data int

	// 启动一个新的 goroutine 来修改 data 的值
	go func() {
		data++
	}()

	// 主 goroutine 等待一段时间，确保 data 的值被修改
	time.Sleep(1 * time.Second) // 这种方式并不优雅，只是为了演示目的

	// 主 goroutine 继续执行，可能会在 data 被修改之前检查它的值
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
