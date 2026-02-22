package main

import "sync"

// Counter 是一个线程安全的计数器，使用互斥锁来保护对 value 的访问。
type Counter struct {
	mu    sync.Mutex // 互斥锁保护 value 的访问
	value int        // 计数器的值
}

// Increment 增加计数器的值
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// GetValue 获取计数器的当前值
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
