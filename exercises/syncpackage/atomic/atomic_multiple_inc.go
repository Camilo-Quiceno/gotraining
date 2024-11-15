package main

import (
	"fmt"         // Import fmt package for formatted I/O operations
	"sync"        // Import sync package for synchronization primitives like WaitGroup
	"sync/atomic" // Import sync/atomic package for atomic operations on shared variables
)

// Global variables
var (
	counter               int64 = 0    // Shared counter variable of type int64
	numberIncreasers      int   = 5    // Number of goroutines that will increment the counter
	numberDecreasers      int   = 3    // Number of goroutines that will decrement the counter
	operationPerGoroutine int64 = 1000 // Number of operations each goroutine will perform
)

func main() {
	wg := sync.WaitGroup{} // Initialize a WaitGroup to wait for all goroutines to finish

	// Start incrementing goroutines
	for i := 0; i < numberIncreasers; i++ {
		wg.Add(1)                // Add a goroutine to the WaitGroup
		go incrementCounter(&wg) // Launch a new goroutine to increment the counter
	}

	// Start decrementing goroutines
	for i := 0; i < numberDecreasers; i++ {
		wg.Add(1)                // Add a goroutine to the WaitGroup
		go decrementCounter(&wg) // Launch a new goroutine to decrement the counter
	}

	wg.Wait() // Wait for all goroutines to complete

	fmt.Println("Final counter value:", counter) // Print the final value of the counter
}

// incrementCounter increments the shared counter by a predefined amount atomically
func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()                                  // Signal that the goroutine is done when the function returns
	atomic.AddInt64(&counter, operationPerGoroutine) // Atomically add the operationPerGoroutine value to the counter
}

// decrementCounter decrements the shared counter by a predefined amount atomically
func decrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()                                   // Signal that the goroutine is done when the function returns
	atomic.AddInt64(&counter, -operationPerGoroutine) // Atomically subtract the operationPerGoroutine value from the counter
}
