package main

import (
	"fmt"
	"runtime"
	"sync"
)

func zoom_num() {
	// 定义一个函数 memConsumed 来测量当前的内存使用情况
	memConsumed := func() uint64 {
		runtime.GC()             // 强制垃圾回收，确保获取到最新的内存使用情况
		var s runtime.MemStats   // 定义一个 runtime.MemStats 结构体来存储内存统计信息
		runtime.ReadMemStats(&s) // 读取当前的内存统计信息并存储在 s 中
		return s.Sys
	}
	// 定义一个只读的通道 c 和一个 WaitGroup wg 来同步 goroutine 的执行
	var c <-chan interface{}
	var wg sync.WaitGroup
	// 定义一个匿名函数 noop，它会在 goroutine 中执行，等待通道 c 的关闭
	noop := func() {
		wg.Done()
		<-c
	}
	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1024)
}

func main() {
	zoom_num()
}
