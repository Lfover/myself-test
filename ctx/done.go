package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(15*time.Second))
	defer cancel() // 确保取消上下文

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("长时间等待")
	case <-ctx.Done():
		fmt.Println("上下文取消")
	}
}
