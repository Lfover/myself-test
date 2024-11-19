package main

import (
	"context"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	//for {
	//	ticker := time.NewTicker()
	//	select {
	//	default:
	//		fmt.Println("hello")
	//	case <-ctx.Done():
	//		return ctx.Err()
	//	}
	//}
	return nil
}

func main1() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}
