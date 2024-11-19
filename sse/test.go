package main

import (
	"fmt"
	"net/http"
	"time"
)

type handle2 struct {
	host string
	port string
}

func main() {
	startServer2()
}
func (this *handle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 设置必要的头部信息
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	for {
		fmt.Fprintf(w, "data:%s", "hello")
		flusher.Flush()
		time.Sleep(1 * time.Second)
	}
}

func startServer2() {
	//被代理的服务器host和port
	h := &handle2{}
	err := http.ListenAndServe(":9003", h)
	if err != nil {

	}
}
