package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ma1in() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}
