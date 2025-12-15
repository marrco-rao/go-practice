package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.elastic.co/apm/v2"
)

func main2() {
	// 初始化APM（以Elastic APM为例）
	apm.DefaultTracer().Close()
	tracer, err := apm.NewTracer("counter-service", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer tracer.Close()

	// 结构化日志配置
	logger := log.New(os.Stdout, "[PROD] ", log.LstdFlags|log.Lmicroseconds|log.LUTC)

	// 业务逻辑
	ctx := context.Background()
	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			// APM事务追踪
			tx := tracer.StartTransaction("IncrementOperation", "background")
			defer tx.End()
			ctx = apm.ContextWithTransaction(ctx, tx)

			// 日志记录
			logger.Printf("Goroutine-%d started | TraceID: %s",
				id, tx.TraceContext().Trace.String())

			defer wg.Done()

			// 业务逻辑
			for j := 0; j < 1000; j++ {
				// APM子跨度
				span := tx.StartSpan("LockOperation", "sync", nil)

				mu.Lock()
				counter++
				mu.Unlock()

				span.End()

				// 关键点采样日志
				if j%200 == 0 {
					logger.Printf("Goroutine-%d progress %d/1000 | %s",
						id, j, time.Now().Format(time.RFC3339Nano))
				}
			}
			logger.Printf("Goroutine-%d completed", id)
		}(i)
	}

	wg.Wait()
	logger.Printf("Final counter: %d | Completed at: %s",
		counter, time.Now().Format(time.RFC3339))
}
