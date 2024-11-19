// package main
//
// import "fmt"
//
//	func main() {
//		ch := make(chan int, 10)
//		ch <- 1
//		ch <- 2
//		ch <- 3
//		close(ch)
//		for i := range ch {
//			fmt.Println(i)
//			fmt.Printf("len(ch): %d, cap(ch): %d\n", len(ch), cap(ch))
//		}
//
// }
package main

import (
	"fmt"
	feature "git.100tal.com/znxx_xpp/feature-ab-client-go"
)

func init() {

}

func main() {
	config := feature.FabConfig{
		Env: feature.TestEnv,
	}
	feature.Register(config, []string{"lite_tts_tal_sn"})
	user := feature.UserWithAttrs(map[string]interface{}{
		"sn": "78PE0823C0100007"})

	pSwitch := feature.BoolValue("lite_tts_tal_sn", user, true)
	if pSwitch {
		fmt.Println("switch on")
	} else {
		fmt.Println("switch off")
	}
}
