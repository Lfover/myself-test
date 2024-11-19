package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 密码，没有则留空
		DB:       0,                // 使用的数据库编号
	})

	pubSub := rdb.Subscribe(ctx, "my_channel") // 订阅频道

	// 处理订阅消息
	ch := pubSub.Channel()
	go func() {
		for msg := range ch {
			fmt.Println("Received message from channel:", msg.Channel, "message:", msg.Payload)
		}
	}()

	// 发布消息
	go func() {
		for {
			err := rdb.Publish(ctx, "my_channel", "Hello, Redis Pub/Sub!").Err()
			if err != nil {
				panic(err)
			}
			// 每隔一秒发布一次消息
			<-time.After(time.Second * 1)
		}
	}()

	// 等待中断信号以优雅地关闭程序
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 取消订阅
	pubSub.Close()
}
