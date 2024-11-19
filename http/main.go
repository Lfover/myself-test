package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport"
	"io/ioutil"
	"net/http"
)

func GetRequestFromCtx(ctx context.Context) (*http.Request, bool) {
	if transPort, ok := transport.FromServerContext(ctx); ok {
		if info, ok := transPort.(trHttp.Transport); ok {
			return info.Request(), true
		}
	}
	return nil, false
}
func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	ctx := context.Background() // 假设这是从某个HTTP处理函数中获取的ctx
	if req, ok := ctx.Value("http.request").(*http.Request); ok {
		// 读取请求体
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println("读取请求体失败:", err)
			return
		}

		fmt.Println("原始请求体:", string(body))
	} else {
		fmt.Println("无法从ctx中获取http.Request对象")
	}
	ctx := context.Background() // 假设这是从某个HTTP处理函数中获取的ctx
	request, success := GetRequestFromCtx(ctx)
	if !success {
		fmt.Println("无法获取HTTP请求对象")
		return
	}

	body, err := readRequestBody(request)
	if err != nil {
		fmt.Println("读取请求体失败:", err)
		return
	}

	fmt.Println("原始请求体:", string(body))
}
