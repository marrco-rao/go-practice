package main

import (
	"bytes"
	"context"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 获取goroutine原始ID（仅调试用）
func getGoID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup, counter *int, mu *sync.Mutex) {
	defer wg.Done()

	// 调试信息
	gid := getGoID()
	log.Printf("[DEBUG] Worker-%d started (GID:%d)", id, gid)
	start := time.Now()

	// 上下文超时控制
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for j := 0; j < 1000; j++ {
		select {
		case <-ctx.Done():
			log.Printf("[WARN] Worker-%d canceled: %v", id, ctx.Err())
			return
		case <-ticker.C:
			log.Printf("[DEBUG] Worker-%d alive (GID:%d, Progress:%d/1000)",
				id, gid, j)
		default:
			mu.Lock()
			*counter++
			mu.Unlock()
		}
	}

	log.Printf("[INFO] Worker-%d completed in %v (GID:%d)",
		id, time.Since(start), gid)
}

func main1() {
	// 开发调试配置
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	// 启动10个worker
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg, &counter, &mu)
	}

	// 监控协程
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				mu.Lock()
				log.Printf("[MONITOR] Current counter: %d", counter)
				mu.Unlock()
			}
		}
	}()

	wg.Wait()
	log.Printf("Final counter: %d", counter)
}
