package main

import "fmt"

func excute() {
	// 启动一个新的 goroutine 来执行 sayHello 函数
	go sayHello()
	// 主 goroutine 继续执行其他代码
	fmt.Println("hello world")
}

func sayGoodbye() {
	go func() {
		fmt.Println("Goodbye, Goroutine!")
	}()
	// 主 goroutine 继续执行其他代码
	fmt.Println("goodbye world")

	// 直接在 goroutine 中定义和调用匿名函数
	sayGo := func() {
		fmt.Println("Go, Goroutine!")
	}
	go sayGo()
}

func sayHello() {
	excute()
	sayGoodbye()
	fmt.Println("Hello, Goroutine!")
}
