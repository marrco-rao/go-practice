package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
锁机制
1. 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
  - 考察点 ： sync.Mutex 的使用、并发数据安全。
*/
var counter int
var mu sync.Mutex

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func lockUse() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("lockUse result:", counter)
}

/*
锁机制
2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
  - 考察点 ：原子操作、并发数据安全。

基本原子操作,type: Int32/Int64/Uint32/Uint64/Pointer/Uintptr
atomic.Add[type](addr *[type],value [type]) (new [type])
atomic.Load[type](addr *[type]) (val [type])
atomic.Store[type](addr *[type],val [type])
atomic.Swap[type](addr *[type],new [type]) (old [type])
atomic.CompareAndSwap[type](addr *[type],old,new [type]) (swapped [type])
*/
func atomicUse() {
	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("atomicUse result:", counter)
}

//func main() {
//	//lockUse()
//	atomicUse()
//
//}
