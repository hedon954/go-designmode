package main

import (
	"fmt"
)

func main() {
	start := Start{}
	p := &patient{Name: "abc"}

	// set the chains
	start.SetNext(&Reception{}).
		SetNext(&DockerCheck{}).
		SetNext(&Payment{}).
		SetNext(&Medicine{})

	// execute
	if err := start.Execute(p); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("success")
}
