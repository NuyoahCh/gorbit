package main

import (
	"sync"
	"sync/atomic"
)

// Context 1: 局部作用域 (函数级上下文)
// 在这个上下文中，变量 i 是局部变量，没有其他 goroutine 可以访问它。
// 因此，从这个函数的视角来看，i++ 是“原子的”，因为没有中间状态会被外部观察到或干扰。
func LocalContext() int {
	var i int
	i++ // 在这个受限的上下文中，它是不可分割的
	return i
}

// Context 2: 共享作用域 (多协程并发上下文)
// 在这个上下文中，i 是一个共享变量。
// 虽然代码看起来和上面一样，但由于并发的存在，i++ 的“检索-修改-存储”三个步骤
// 可能会被其他协程中断，导致竞态条件（Race Condition）。
func SharedContext() int {
	var i int
	var wg sync.WaitGroup

	// 启动 1000 个协程同时执行 i++
	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 这里的 i++ 不是原子的！
			// 多个协程可能同时检索到相同的旧值，增加后存回，导致覆盖掉其他协程的修改。
			i++
		}()
	}
	wg.Wait()
	return i // 结果通常会小于 1000
}

// Context 3: 显式的原子操作上下文 (机器/指令级上下文)
// 如果我们需要在并发环境下保持原子性，必须改变操作的方式，
// 使用底层硬件提供的原子指令。
func AtomicPackageContext() int64 {
	var i int64
	var wg sync.WaitGroup

	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 即使在多协程并发的上下文中，这个操作也是原子的。
			// 它通过底层的 CPU 指令锁（如 LOCK 前缀）保证了操作的不可分割性。
			atomic.AddInt64(&i, 1)
		}()
	}
	wg.Wait()
	return i // 结果永远是 1000
}
