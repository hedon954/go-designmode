package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a decision tree
	tree := buildDecisionTree()

	// Main game loop
	var i int
	for {
		i++
		time.Sleep(3 * time.Second)
		if tree.Evaluate() {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
