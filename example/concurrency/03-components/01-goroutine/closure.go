package main

import (
	"fmt"
	"sync"
)

func example01() {
	var wg sync.WaitGroup
	// 外部函数中的变量
	salutation := "Hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		// goroutine 修改了变量 salutation 的值
		// 这里的 salutation 是外部函数中的变量，goroutine 内部的匿名函数可以访问和修改它
		salutation = "Welcome"
	}()
	wg.Wait()
	fmt.Println(salutation) // 输出：Welcome
}

func example02() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			// 这里引用字符串类型的切片作为创建循环变量 salutation 的值
			fmt.Println(salutation)
		}(salutation) // 将循环变量作为参数传递给匿名函数，避免闭包捕获问题
	}
	// 等待所有 goroutine 完成
	wg.Wait()
}

func main() {
	example01()
	example02()
}
