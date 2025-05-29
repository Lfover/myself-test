package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

//
//LockWithTimeout()
//Unlock

/*
给定一个正整数N，请统计所有小于等于N的正整数中，
满足其二进制表示中1的个数为质数的数字个数。
*/
func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		// v已经是int类型
		fmt.Printf("整数: %d (类型: %T)\n", v, v)
	case string:
		// v已经是string类型
		fmt.Printf("字符串: %s (类型: %T)\n", v, v)
	case []byte:
		// v已经是[]byte类型
		fmt.Printf("字节切片: %s (类型: %T)\n", v, v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}
func main() {
	typeSwitch("1")
	value6 := interface{}(byte(127))
	switch t := value6.(type) {
	case uint16:
		fmt.Println("uint8 or uint16")
	case byte:
		fmt.Printf("byte")
	default:
		fmt.Printf("unsupported type: %T", t)
	}
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
				panic("1")
			}
			time.Sleep(time.Microsecond)
		}

	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)

	}
	time.Sleep(time.Second * 5)
}
func TestA(test *testing.T) {

}

// 1122333
func GetCount(n, x int) int {
	s := cast.ToString(n)
	x1 := cast.ToString(x)

	s1 := strings.Count(s, x1)
	fmt.Println(s1)
	count := 0
	for n > 0 {
		if n%10 == x {
			count++
		}
		n /= 10
	}
	return count

}

func LockWithTimeout() {

}
