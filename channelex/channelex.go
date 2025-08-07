package main

import (
	"flag"
	"fmt"
	"sync"
)

// --- WaitGroup Implementation ---

func generateNumbersWg(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()
	return ch
}

func consumeNumbersWg(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Received (WaitGroup):", num)
	}
}

// --- Done Channel Implementation ---

func generateNumbersChan(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()
	return ch
}

func consumeNumbersChan(ch <-chan int, done chan<- bool) {
	defer func() {
		done <- true
	}()
	for num := range ch {
		fmt.Println("Received (Done Chan):", num)
	}
}

func main() {
	syncMethod := flag.String("sync", "wg", "Synchronization method: 'wg' for WaitGroup, 'chan' for done channel.")
	flag.Parse()

	switch *syncMethod {
	case "wg":
		fmt.Println("Using sync.WaitGroup for synchronization.")
		var wg sync.WaitGroup
		numbersCh := generateNumbersWg(10)
		wg.Add(1)
		go consumeNumbersWg(numbersCh, &wg)
		wg.Wait()
		fmt.Println("WaitGroup finished. Exiting.")

	case "chan":
		fmt.Println("Using a done channel for synchronization.")
		done := make(chan bool)
		numbersCh := generateNumbersChan(10)
		go consumeNumbersChan(numbersCh, done)
		<-done
		fmt.Println("Done channel received signal. Exiting.")

	default:
		fmt.Println("Invalid sync method. Please use 'wg' or 'chan'.")
	}
}
