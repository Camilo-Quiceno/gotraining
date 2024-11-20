package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Define a list of predefined tasks to be processed.
	tasks := []string{
		"Task 1: Data Processing",
		"Task 2: File Download",
		"Task 3: Image Rendering",
		"Task 4: Report Generation",
		"Task 5: Email Sending",
	}

	// Create channels for task distribution and result collection.
	in := make(chan string)  // Channel to send tasks to workers.
	out := make(chan string) // Channel to receive completed tasks from workers.

	var wg sync.WaitGroup // WaitGroup to synchronize worker goroutines.

	done := make(chan struct{}) // Channel to signal cancellation to workers.

	numWorkers := 3 // Number of worker goroutines to launch.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)                     // Increment WaitGroup counter for each worker.
		go worker(in, done, out, &wg) // Launch worker goroutine.
	}

	// Goroutine to trigger cancellation after 1 second.
	go cancellation(done)

	// Goroutine to send tasks to the 'in' channel.
	go func() {
		for _, task := range tasks {
			select {
			case in <- task:
				// Successfully sent task to worker.
			case <-done:
				// Cancellation signal received; stop sending tasks.
				close(in) // Close 'in' channel to notify workers no more tasks will be sent.
				return    // Exit the goroutine.
			}
		}
		close(in) // Close 'in' channel after sending all tasks.
	}()

	// Goroutine to close the 'out' channel after all workers have finished.
	go func() {
		wg.Wait()  // Wait for all workers to finish.
		close(out) // Close 'out' channel to signal that no more results will be sent.
	}()

	// Collect results from the 'out' channel.
	numTasks := 0 // Counter for completed tasks.
	for range out {
		numTasks++ // Increment counter for each completed task received.
	}

	// Display the summary of completed and cancelled tasks.
	fmt.Printf("Total tasks completed: %d\n", numTasks)
	fmt.Printf("Total tasks cancelled: %d\n", len(tasks)-numTasks)
}

// worker is a function that processes tasks from the 'tasks' channel.
// It listens for cancellation signals from the 'done' channel.
// Completed tasks are sent to the 'out' channel.
// The WaitGroup is decremented when the worker exits.
func worker(tasks <-chan string, done <-chan struct{}, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement WaitGroup counter when the worker exits.

	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				// 'tasks' channel is closed; no more tasks to process.
				return // Exit the worker goroutine.
			}

			fmt.Printf("Starting %s\n", task) // Log task start.

			// Simulate task processing by sleeping for a random duration up to 4 seconds.
			randomDuration := time.Duration(rand.Intn(4000)) * time.Millisecond
			time.Sleep(randomDuration) // Simulate work.

			select {
			case out <- task:
				// Successfully sent completed task to 'out' channel.
				fmt.Printf("%s Completed\n", task) // Log task completion.
			case <-done:
				// Cancellation signal received while trying to send the result.
				return // Exit the worker goroutine.
			}
		case <-done:
			// Cancellation signal received; exit the worker.
			return
		}
	}
}

// cancellation is a function that signals cancellation by closing the 'done' channel after a delay.
func cancellation(done chan struct{}) {
	time.Sleep(1 * time.Second) // Wait for 1 second before triggering cancellation.
	close(done)                 // Close the 'done' channel to signal cancellation to workers.
}
