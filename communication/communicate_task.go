package communication

// You have a producer of data (e.g., reading from a file, fetching from an API) and one or more consumers that process this data.
// The producer and consumers operate at different speeds.

import (
	"fmt"
	"time"
)

// producer sends numbers to a channel
func producer(dataCh chan<- int, num int) {
	for i := 0; i < num; i++ {
		fmt.Printf("Producing %d\n", i)
		dataCh <- i                        // Send data to the channel
		time.Sleep(time.Millisecond * 100) // Simulate work
	}
	close(dataCh) // Close the channel when done sending
	fmt.Println("Producer finished")
}

// consumer receives numbers from a channel and processes them
func consumer(dataCh <-chan int, done chan<- bool, id int) {
	fmt.Printf("Consumer %d started\n", id)
	for data := range dataCh { // Receive data from the channel
		fmt.Printf("Consumer %d processing %d\n", id, data)
		time.Sleep(time.Millisecond * 300) // Simulate work
	}
	fmt.Printf("Consumer %d finished\n", id)
	done <- true // Signal completion
}

// RunCommunicateTask sets up and runs the producer-consumer simulation.
func RunCommunicateTask() {
	dataChannel := make(chan int, 5)  // Buffered channel with capacity 5
	doneChannel := make(chan bool, 2) // Channel to signal consumer completion

	// Start the producer goroutine
	go producer(dataChannel, 10)

	// Start consumer goroutines
	go consumer(dataChannel, doneChannel, 1)
	go consumer(dataChannel, doneChannel, 2)

	// Wait for both consumers to finish
	<-doneChannel
	<-doneChannel

	fmt.Println("Producer and consumers finished.")
}
