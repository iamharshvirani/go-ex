package processing

// A basic in-memory task processing system can consist of a channel to act as a task queue
// and a pool of worker goroutines that read tasks from the channel and execute them.

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work (e.g., a function to execute)
type Task func()

// TaskProcessor is a simple in-memory task queue and worker pool
type TaskProcessor struct {
	tasks chan Task      // Channel to queue tasks
	wg    sync.WaitGroup // To wait for workers to finish
}

// NewTaskProcessor creates a new TaskProcessor with a given number of workers
func NewTaskProcessor(numberOfWorkers int) *TaskProcessor {
	processor := &TaskProcessor{
		tasks: make(chan Task), // Unbuffered channel for simplicity initially
	}

	// Start the worker goroutines
	for i := 0; i < numberOfWorkers; i++ {
		processor.wg.Add(1)
		go processor.worker(i)
	}

	return processor
}

// worker is a goroutine that processes tasks from the tasks channel
func (p *TaskProcessor) worker(id int) {
	defer p.wg.Done()
	fmt.Printf("Worker %d started.\n", id)
	for task := range p.tasks { // Read tasks from the channel
		fmt.Printf("Worker %d processing task...\n", id)
		task() // Execute the task
		fmt.Printf("Worker %d finished task.\n", id)
	}
	fmt.Printf("Worker %d shutting down.\n", id)
}

// Submit adds a task to the queue
func (p *TaskProcessor) Submit(task Task) {
	p.tasks <- task // Send the task to the channel
}

// Stop closes the task channel and waits for workers to finish
func (p *TaskProcessor) Stop() {
	close(p.tasks) // Signal that no more tasks will be sent
	p.wg.Wait()    // Wait for all workers to finish
	fmt.Println("Task processor stopped.")
}

func RunTaskProcessor() {
	// Create a task processor with 3 workers
	processor := NewTaskProcessor(3)

	// Submit some tasks
	for i := 1; i <= 10; i++ {
		taskID := i
		processor.Submit(func() {
			fmt.Printf("Executing task %d\n", taskID)
			time.Sleep(time.Millisecond * 500) // Simulate work
		})
	}

	// In a real application, you'd typically stop the processor on shutdown signal
	// For this example, let's wait a bit and then stop
	time.Sleep(3 * time.Second) // Allow some tasks to be processed
	processor.Stop()
}
