package processing

// You need to process a batch of data, such as sending notifications to a list of users, processing images, or performing independent calculations.
// Doing these tasks sequentially would be slow.

import (
	"fmt"
	"sync"
	"time"
)

// processDataItem simulates processing a single data item.
func processDataItem(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes

	fmt.Printf("Processing item %d...\n", id)
	time.Sleep(time.Millisecond * time.Duration(500+id*50)) // Simulate variable processing time
	fmt.Printf("Finished processing item %d.\n", id)
}

// RunProcessItems sets up and processes a batch of data items concurrently.
func RunProcessItems() {
	dataItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup // Used to wait for a collection of goroutines to finish

	fmt.Println("Starting concurrent processing...")

	for _, item := range dataItems {
		wg.Add(1)                     // Increment the counter for each goroutine
		go processDataItem(item, &wg) // Launch a goroutine for each item
	}

	wg.Wait() // Block until the counter becomes zero

	fmt.Println("All items processed.")
}
