package hungrygophers

import (
	"fmt"
	"sync"
	"time"
)

// The Gopher struct and main function setup remain the same.
// We just change the `eat` method and add a "table" semaphore.

type GopherWithSemaphore struct {
	id        int
	leftFork  *sync.Mutex
	rightFork *sync.Mutex
	eatCount  int
}

// table is a semaphore that allows at most 4 Gophers to "sit down"
var table = make(chan struct{}, 4)

func (g *GopherWithSemaphore) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Printf("Gopher %d is thinking.\n", g.id)
		time.Sleep(time.Millisecond * time.Duration(100+g.id*50))

		fmt.Printf("Gopher %d is hungry and wants to sit at the table.\n", g.id)
		table <- struct{}{} // Acquire a spot at the table (semaphore)
		fmt.Printf("Gopher %d sat at the table.\n", g.id)

		// Now that we are at the table, we can try to pick up forks
		g.leftFork.Lock()
		fmt.Printf("Gopher %d picked up left fork.\n", g.id)

		g.rightFork.Lock()
		fmt.Printf("Gopher %d picked up right fork.\n", g.id)

		fmt.Printf("Gopher %d is eating.\n", g.id)
		g.eatCount++
		time.Sleep(time.Millisecond * 100)

		g.rightFork.Unlock()
		g.leftFork.Unlock()
		fmt.Printf("Gopher %d put down forks and is leaving the table.\n", g.id)

		<-table // Release the spot at the table
	}
}

func RunGopherSemaphore() {
	// Setup is identical to the first example...
	numGophers := 5
	forks := make([]*sync.Mutex, numGophers)
	for i := 0; i < numGophers; i++ {
		forks[i] = &sync.Mutex{}
	}

	gophers := make([]*GopherWithSemaphore, numGophers)
	for i := 0; i < numGophers; i++ {
		gophers[i] = &GopherWithSemaphore{
			id:        i,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numGophers],
		}
	}

	var wg sync.WaitGroup
	wg.Add(numGophers)

	fmt.Println("Dinner is starting!")
	startTime := time.Now()

	for i := 0; i < numGophers; i++ {
		go gophers[i].eat(&wg)
	}

	wg.Wait()
	fmt.Printf("\nDinner is over after %v.\n", time.Since(startTime))
	for _, g := range gophers {
		fmt.Printf("Gopher %d ate %d times.\n", g.id, g.eatCount)
	}
}
