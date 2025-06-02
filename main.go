package main

import "fmt"

func main() {
	// Hardcoded variable to choose the program to run
	// Options: "communicate", "process"
	programToRun := "communicate"

	switch programToRun {
	case "communicate":
		fmt.Println("Running Communicate Task Program...")
		runCommunicateTask()
	case "process":
		fmt.Println("Running Process Items Program...")
		runProcessItems()
	default:
		fmt.Println("Invalid program choice. Please choose 'communicate' or 'process'.")
	}
}
