package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

//func main() {
//	extra := map[string]interface{}{}
//	var str []string
//	if "一年级" != "" {
//		str = append(str, "一年级")
//	}
//	if "比配" != "" {
//		str = append(str, "比配")
//	}
//	if len(str) > 0 {
//		Tags := fmt.Sprintf("[\"%s\"]", strings.Join(str, "\",\""))
//		fmt.Println(Tags)
//		extra["tags"] = Tags
//	}
//	if "rec.Author" != "" {
//		extra["作者："] = "rec.Author"
//	}
//	fmt.Println(str, extra)
//}

func main() {
	arr := []string{"一年级", "比配"}
	strArr, _ := json.Marshal(arr)
	fmt.Println(string(strArr))
	limiter := rate.NewLimiter(100, 1)
	for {
		go func() {}()
		time.Sleep(limiter.Reserve().Delay())
	}
}
