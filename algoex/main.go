package main

import (
	"fmt"
	"go-ex/algoex/algos"
)

func main() {
	programToRun := "lru" // You can change this to "process" to test the other part

	switch programToRun {
	case "mergeelemets":
		fmt.Println("Merge Elements Program...")
		algos.RunMergeElements()
	case "lru":
		fmt.Println("LRU Cache Program...")
		algos.RunLRUCache()
	default:
		fmt.Println("Invalid program choice. Please choose a valid program")
	}
}
