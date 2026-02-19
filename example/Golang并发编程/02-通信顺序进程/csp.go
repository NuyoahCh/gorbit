package main

import (
	"fmt"
	"time"
)

// 示例1: 基础的CSP通信
// 演示两个goroutine通过channel进行通信
func basicCSP() {
	fmt.Println("=== 基础CSP通信示例 ===")

	// 创建一个字符串channel
	messages := make(chan string)

	// 启动第一个goroutine - 发送者
	go func() {
		time.Sleep(time.Second)
		messages <- "hello"
	}()

	// 启动第二个goroutine - 发送者
	go func() {
		time.Sleep(time.Second * 2)
		messages <- "world"
	}()

	// 主goroutine - 接收者
	for i := 0; i < 2; i++ {
		msg := <-messages
		fmt.Println("收到消息:", msg)
	}
	fmt.Println()
}

// 示例2: 工作池模式
// 多个worker从channel接收任务
func workerPoolCSP() {
	fmt.Println("=== 工作池模式 ===")

	// 创建任务和结果channel
	jobs := make(chan int, 5)
	results := make(chan string, 5)

	// 启动3个worker
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		go func(workerID int) {
			for job := range jobs {
				fmt.Printf("Worker %d 开始处理任务 %d\n", workerID, job)
				time.Sleep(time.Millisecond * 500)
				results <- fmt.Sprintf("Worker %d 完成任务 %d", workerID, job)
			}
		}(w)
	}

	// 发送任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for a := 1; a <= 5; a++ {
		fmt.Println("结果:", <-results)
	}
	fmt.Println()
}

// 示例3: 扇出/扇入模式
// 一个发送者，多个接收者（扇出）
// 多个发送者，一个接收者（扇入）
func fanOutFanInCSP() {
	fmt.Println("=== 扇出/扇入模式 ===")

	// 创建channel
	numbers := make(chan int)
	squares := make(chan int)
	cubes := make(chan int)

	// 扇出: 数据发送到多个处理器
	go func() {
		for i := 1; i <= 3; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// 处理器1: 计算平方
	go func() {
		for n := range numbers {
			squares <- n * n
		}
		close(squares)
	}()

	// 处理器2: 计算立方
	go func() {
		for n := range numbers {
			cubes <- n * n * n
		}
		close(cubes)
	}()

	// 扇入: 合并多个channel的数据
	for i := 0; i < 3; i++ {
		fmt.Printf("数字: %d, 平方: %d, 立方: %d\n", i+1, <-squares, <-cubes)
	}
	fmt.Println()
}

// 示例4: select语句 - 多channel复用
// 同时监听多个channel
func selectCSP() {
	fmt.Println("=== Select多路复用 ===")

	tick := time.Tick(time.Millisecond * 100)
	boom := time.After(time.Millisecond * 500)

	for {
		select {
		case <-tick:
			fmt.Println("嘀... (每100ms)")
		case <-boom:
			fmt.Println("轰! (500ms后触发)")
			return
		}
	}
}

// 示例5: Pipeline模式
// 多个阶段通过channel串联
func pipelineCSP() {
	fmt.Println("=== Pipeline管道模式 ===")

	// 第一阶段: 生成数字
	generate := func() <-chan int {
		out := make(chan int)
		go func() {
			for i := 1; i <= 5; i++ {
				out <- i
			}
			close(out)
		}()
		return out
	}

	// 第二阶段: 数字平方
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// 第三阶段: 打印结果
	for n := range square(square(generate())) {
		fmt.Printf("最终结果: %d\n", n)
	}
	fmt.Println()
}

// 示例6: 超时控制
// 使用select实现超时机制
func timeoutCSP() {
	fmt.Println("=== 超时控制示例 ===")

	resultChan := make(chan string)

	// 模拟可能超时的操作
	go func() {
		time.Sleep(time.Millisecond * 800)
		resultChan <- "操作完成"
	}()

	// 设置500ms超时
	select {
	case res := <-resultChan:
		fmt.Println("收到结果:", res)
	case <-time.After(time.Millisecond * 500):
		fmt.Println("操作超时")
	}
	fmt.Println()
}

func main() {
	basicCSP()
	workerPoolCSP()
	fanOutFanInCSP()
	selectCSP()
	pipelineCSP()
	timeoutCSP()
}
