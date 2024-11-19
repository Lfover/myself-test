package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main3() {
	p := Person{Name: "Alice", Age: 20}
	fmt.Println(p)
	bytes, _ := json.Marshal(p)
	fmt.Println("Byte count:", len(bytes))
	fmt.Println("Character count:", len(string(bytes)))
}
func main() {
	s := []string{"a", "b", "c"}
	for i, v := range s {
		go func() {
			fmt.Println(i, v)
		}()

	}
	time.Sleep(1 * time.Second)
}
