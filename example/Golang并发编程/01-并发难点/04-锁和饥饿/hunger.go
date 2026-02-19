package main

import (
	"fmt"
	"sync"
	"time"
)

func example01() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	// 贪婪工人：持续尝试获取锁并执行工作
	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()               // 贪婪地获取锁
			time.Sleep(3 * time.Nanosecond) // 模拟工作
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker count: %d\n", count)
	}

	// 定期工人：定期获取锁，打乱贪婪工人的节奏
	populateWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()               // 定期获取锁，打乱贪婪工人的节奏
			time.Sleep(1 * time.Nanosecond) // 模拟工作
			sharedLock.Unlock()

			sharedLock.Lock()               // 定期获取锁，打乱贪婪工人的节奏
			time.Sleep(1 * time.Nanosecond) // 模拟工作
			sharedLock.Unlock()

			sharedLock.Lock()               // 定期获取锁，打乱贪婪工人的节奏
			time.Sleep(1 * time.Nanosecond) // 模拟工作
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Populate worker count: %d\n", count)
	}
	wg.Add(2)
	go greedyWorker()
	go populateWorker()
	wg.Wait()
}
