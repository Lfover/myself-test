//
//import (
//	"context"
//	"net/http"
//	"time"
//
//	"github.com/go-kratos/kratos/v2/log"
//	"github.com/go-kratos/kratos/v2/transport/http"
//)
//
//// EventServer is a HTTP server that sends server-sent events.
//type EventServer struct {
//	logger *log.Helper
//}
//
//// NewEventServer creates a new EventServer instance.
//func NewEventServer(logger log.Logger) *EventServer {
//	return &EventServer{
//		logger: log.NewHelper(logger),
//	}
//}
//
//// ServeHTTP implements the http.Handler interface.
//func (s *EventServer) ServeHTTP(ctx context.Context, req *http.Request) (interface{}, error) {
//	// 设置响应头
//	resp.Header().Set("Content-Type", "text/event-stream")
//	resp.Header().Set("Cache-Control", "no-cache")
//	resp.Header().Set("Connection", "keep-alive")
//
//	// 创建一个 channel 来控制发送流程
//	closeNotify := req.Context().Done()
//
//	// 使用select来监听多个channel
//	for {
//		select {
//		case <-closeNotify:
//			s.logger.Infof("Client has disconnected")
//			return nil, nil
//		default:
//			// 发送事件
//			now := time.Now().Format(time.RFC3339)
//			resp.BodyWriter().Write([]byte("data: The current time is " + now + "\n\n"))
//			resp.BodyWriter().Flush()
//		}
//
//		// 每隔一段时间发送一次事件
//		time.Sleep(1 * time.Second)
//	}
//}
//
//// RegisterHTTP registers the EventServer as an HTTP handler.
//func (s *EventServer) RegisterHTTP(server *http.Server) {
//	server.Handle("/events", http.HandlerFunc(s.ServeHTTP))
//}

package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	client := &http.Client{}
	s := `{
    "message": [
        {
            "type": 1,
            "msg": "写一个冒泡排序",
            "current": 1
        }
    ]
}`
	req, err := http.NewRequest("POST", "http://hmi.chengjiukehu.com/paas-proxy/sse/send", strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := fmt.Sprintf("%x%d", md5.Sum([]byte(timestamp)), rand.Intn(100000))
	appid := "200020"
	secret := "pkoi4mfgryc3puphpysqxr6spe2cpd30"
	deviceId := "1234wadwadd567"
	version := "1.0.0"
	osVersion := "T202210"
	platform := "iOS"

	headerStr := secret + "&X-Genie-Timestamp=" + timestamp + "&X-Genie-Nonce=" + nonce + "&X-Genie-DeviceId=" + deviceId
	sign := fmt.Sprintf("%x", md5.Sum([]byte(headerStr)))

	req.Header.Set("X-Genie-DeviceId", deviceId)
	req.Header.Set("X-Genie-Timestamp", timestamp)
	req.Header.Set("X-Genie-Nonce", nonce)
	req.Header.Set("X-Genie-Sign", sign)
	req.Header.Set("X-Genie-AppId", appid)
	req.Header.Set("X-Genie-Version", version)
	req.Header.Set("X-Genie-Osversion", osVersion)
	req.Header.Set("X-Genie-Platform", platform)
	req.Header.Set("X-Genie-TraceId", "123")
	req.Header.Set("baggage", "origin=SseBaiduChat")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
	//for {
	//	buf := make([]byte, 1024)
	//	n, err := resp.Body.Read(buf)
	//	if n > 0 {
	//		fmt.Println(string(buf[:n]))
	//	}
	//	if err != nil {
	//		break
	//	}
	//}

}
