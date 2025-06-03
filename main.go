package main

import (
	"fmt"
	"go-ex/communication"
	"go-ex/customcache"
	"go-ex/pkg/sharedresource"
	"go-ex/processing"
	"go-ex/ratelimiter"
)

func main() {
	// Hardcoded variable to choose the program to run
	// Options: "communicate", "process", "sharedresource", "sharedresourcemap", "simplecache", "concurrentcache", "ratelimiter", "taskprocessor"
	programToRun := "taskprocessor" // You can change this to "process" to test the other part

	switch programToRun {
	case "communicate":
		fmt.Println("Running Communicate Task Program...")
		communication.RunCommunicateTask()
	case "process":
		fmt.Println("Running Process Items Program...")
		processing.RunProcessItems()
	case "sharedresource":
		fmt.Println("Running Shared Resource Program...")
		sharedresource.RunSharedResource()
	case "sharedresourcemap":
		fmt.Println("Running Shared Resource Map Program...")
		sharedresource.RunSharedResourceMap()
	case "simplecache":
		fmt.Println("Running Simple Cache Program...")
		customcache.RunSimpleCache()
	case "concurrentcache":
		fmt.Println("Running Concurrent Cache Program...")
		customcache.RunConcurrentCache()
	case "ratelimiter":
		fmt.Println("Running Rate Limiter Program...")
		ratelimiter.RunRateLimiter()
	case "taskprocessor":
		fmt.Println("Running Task Processor Program...")
		processing.RunTaskProcessor()
	default:
		fmt.Println("Invalid program choice. Please choose a valid program")
	}
}
