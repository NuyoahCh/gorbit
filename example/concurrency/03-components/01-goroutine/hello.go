package main

import (
	"fmt"
	"sync"
)

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

func forkModel1() {
	// 创建一个匿名函数，并将其作为一个 goroutine 启动
	sayHello := func() {
		fmt.Println("Hello")
	}
	// 启动一个新的 goroutine 来执行 sayHello 函数
	go sayHello()
	// 继续执行自己的逻辑
	fmt.Println("World")
}

func forkModel2() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done() // 在 goroutine 完成时调用 Done 方法
		fmt.Println("Hello")
	}
	wg.Add(1)
	go sayHello()
	// 这里就是连接点使用的方式
	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("World")
}

func sayHello() {
	// 模拟一些工作，目的为了减少编译器中的 Warnings
	excute()
	sayGoodbye()
	forkModel1()
	forkModel2()
	fmt.Println("Hello, Goroutine!")
}
