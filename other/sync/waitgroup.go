package main

import (
	"fmt"
	"sync"
	"time"
)

/*
基本概念

	sync.WaitGroup 通过计数器机制实现等待功能：

	Add(delta int)：增加或减少等待的 goroutine 数量
	Done()：相当于 Add(-1)，表示一个 goroutine 已完成
	Wait()：阻塞直到计数器归零
*/
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 为每个goroutine增加计数器
		go worker(i, &wg)
	}

	wg.Wait() // 等待所有goroutine完成
	fmt.Println("所有工作已完成")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 函数结束时减少计数器

	fmt.Printf("工人 %d 开始工作\n", id)
	time.Sleep(time.Second * time.Duration(id))
	fmt.Printf("工人 %d 完成工作\n", id)
}
