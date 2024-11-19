package main

import (
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	// 初始化 enforcer
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("Error in Casbin NewEnforcer: %v", err)
	}

	// 检查权限
	if res, err := e.Enforce("alice", "data1", "read"); err == nil && res {
		log.Println("Alice can read data1")
	} else {
		log.Println("Alice can't read data1")
	}
}
