package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Hello World")
	}()
	go func() {

		defer wg.Done()
		fmt.Println("Hello World2")
	}()
	wg.Wait()
	fmt.Println("Done")
	return
}
