package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "O2EVb1fwE4AtPyi1oswnDw=="
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
}
