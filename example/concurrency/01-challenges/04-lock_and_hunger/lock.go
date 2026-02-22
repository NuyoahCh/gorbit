package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// value 结构体包含一个互斥锁和一个整数值
type value struct {
	mu    sync.Mutex // 互斥锁，用于保护 value 字段的访问
	value int        // 整数值
}

func deadlock() {
	// 创建一个 WaitGroup 来等待所有 goroutine 完成
	var wg sync.WaitGroup
	// 定义一个函数来计算两个 value 的和，并使用互斥锁来保护访问
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		// 锁定 v1，确保在访问 v1.value 时不会被其他 goroutine 干扰
		v1.mu.Lock() // 锁定 v1
		defer v1.mu.Unlock()
		// 模拟一些工作，增加死锁发生的可能性
		time.Sleep(2 * time.Second)
		// 锁定 v2，确保在访问 v2.value 时不会被其他 goroutine 干扰
		v2.mu.Lock()
		defer v2.mu.Unlock()
		// 计算 v1 和 v2 的和，并打印结果
		fmt.Printf("sum: %v\n", v1.value+v2.value)
	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b) // 计算 a 和 b 的和
	go printSum(&b, &a) // 计算 b 和 a 的和
	wg.Wait()
}

func livelock() {
	// 创建 cadence 条件变量，用于协调 goroutine 之间的执行
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast() // 广播信号，唤醒所有等待的 goroutine
		}
	}()
	takeStep := func() {
		cadence.L.Lock()   // 锁定 cadence 条件变量的互斥锁
		cadence.Wait()     // 等待 cadence 条件变量的信号
		cadence.L.Unlock() // 解锁 cadence 条件变量的互斥锁
	}
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, "Trying to step %s\n", dirName) // 输出尝试的方向
		atomic.AddInt32(dir, 1)                          // 增加 dir 的值，表示尝试这个方向
		takeStep()                                       // 等待 cadence 条件变量的信号
		if atomic.LoadInt32(dir) == 1 {                  // 检查 dir 的值是否为 1，表示成功尝试这个方向
			fmt.Fprintf(out, "Stepped %s\n", dirName) // 输出成功的方向
			return true
		}
		takeStep()               // 等待 cadence 条件变量的信号
		atomic.AddInt32(dir, -1) // 减少 dir 的值，表示放弃这个方向
		return false
	}
	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out) // 尝试向左移动
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out) // 尝试向右移动
	}
	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() {
			fmt.Printf("%s:\n%s", name, out.String()) // 输出 goroutine 的名称和尝试的结果
		}()
		defer walking.Done()
		fmt.Fprintf(&out, "Starting to walk %s\n", name) // 输出开始行走的消息
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) { // 尝试向左或向右移动
				return // 如果成功移动，退出函数
			}
		}
		fmt.Fprintf(&out, "Couldn't walk %s\n", name) // 输出无法行走的消息
	}
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "person 1") // 启动第一个 goroutine 来行走
	go walk(&peopleInHallway, "person 2") // 启动第二个 goroutine 来行走
	peopleInHallway.Wait()                // 等待所有 goroutine 完成
}
