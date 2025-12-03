package main

import (
	"fmt"
	"math/rand"
	"sync"
	//"time"
)

/*
Channel
1. 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
  - 考察点 ：通道的基本使用、协程间通信。
*/
func verifyRoutineCommunication() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func(max int) {
		defer wg.Done()
		for i := 1; i <= max; i++ {
			fmt.Printf("@生产者发送: %d\n", i)
			ch <- int(i)
		}
		close(ch) // 关键：发送完毕关闭通道
	}(10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range ch {
			fmt.Printf("消费者接收: %d\n", n)
		}
	}()
	wg.Wait()
	fmt.Println("finish")
}

/*
2. 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
   - 考察点 ：通道的缓冲机制。
*/

func verifyRoutineCommunication1() {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(num int) {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			tmp := rand.Intn(1000)
			ch <- tmp
			fmt.Printf("@生产者发送: %d (缓冲剩余%d/%d)\n",
				tmp, len(ch), cap(ch))
		}
		close(ch) // 关键：发送完毕关闭通道
	}(100)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range ch {
			fmt.Printf("消费者接收: %d (缓冲剩余%d/%d)\n", n, len(ch), cap(ch))
		}
	}()
	wg.Wait()
	fmt.Println("finish")
}

//func main() {
//	//verifyRoutineCommunication()
//	verifyRoutineCommunication1()
//}
