package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/*
1,题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func routineWork(o int, e int) {
	var wg sync.WaitGroup
	fmt.Println("start")
	wg.Add(1)
	go func(o int) {
		defer wg.Done()
		if o < 0 {
			return
		}
		odd := []int{}
		for i := 1; i <= o; i++ {
			if i%2 == 1 {
				odd = append(odd, i)
			}
		}
		fmt.Println("odd:", odd)
	}(o)
	wg.Add(1)
	go func(e int) {
		defer wg.Done()
		if e < 0 {
			return
		}
		even := []int{}
		for i := 1; i <= e; i++ {
			if i%2 == 0 {
				even = append(even, i)
			}
		}
		fmt.Println("even:", even)
	}(e)
	wg.Wait()
	fmt.Println("finish")
}

func ex2_test1() {
	routineWork(10, 10)
}

/*
2.题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
type TaskResult struct {
	name     string
	Duration time.Duration
}

// 超时，Panic，任务报错 等内容可以做扩展
func measureTaskTime(tasks []func()) {
	var wg sync.WaitGroup
	durations := make([]TaskResult, len(tasks))
	s := time.Now()
	for i, task := range tasks {
		wg.Add(1)
		go func(idx int, task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			taskName := "task-" + strconv.Itoa(idx)
			durations[idx] = TaskResult{taskName, time.Since(start)}
		}(i, task)
	}
	wg.Wait()
	fmt.Println("execution time:", time.Since(s))
	fmt.Println("durations:", durations)
	fmt.Println("finish")
}
func ex2_test2() {
	var tasks []func()
	for i := 1; i <= 10; i++ {
		tasks = append(tasks, func() {
			du := rand.Intn(10)
			fmt.Println("Exec task_", i, ":", du)
			time.Sleep(time.Duration(du) * time.Second)
		})
	}
	measureTaskTime(tasks)
}

//func main() {
//	//ex2_test1()
//	ex2_test2()
//}
