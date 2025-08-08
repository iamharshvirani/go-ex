package main

import (
	"fmt"
	"go-ex/despatterns/patterns"
)

func main() {
	// Change this variable to run a different pattern
	// Options: "singleton", "constructor", "adapter", "factory", "decorator", "observer"
	patternToRun := "decorator" // Changed to "decorator" to test the new pattern

	switch patternToRun {
	case "singleton":
		patterns.RunSingleton()
	case "constructor":
		patterns.RunConstructor()
	case "adapter":
		patterns.RunAdapter()
	case "factory":
		patterns.RunFactory()
	case "decorator":
		patterns.RunDecorator()
	case "observer":
		patterns.RunObserver()
	default:
		fmt.Println("Invalid pattern specified")
	}
}
