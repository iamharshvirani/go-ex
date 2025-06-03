package main

import (
	"fmt"
	"go-ex/communication"
	"go-ex/pkg/sharedresource"
	"go-ex/processing"
)

func main() {
	// Hardcoded variable to choose the program to run
	// Options: "communicate", "process"
	programToRun := "communicate" // You can change this to "process" to test the other part

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
	default:
		fmt.Println("Invalid program choice. Please choose a valid program")
	}
}
