package main

import (
	"fmt"         // Import fmt package for formatted I/O operations
	"sync"        // Import sync package for synchronization primitives like WaitGroup
	"sync/atomic" // Import sync/atomic package for atomic operations on shared variables
)

// Global variables
var (
	counter       int64 = 0    // Shared counter variable of type int64
	numIncrements int64 = 1000 // Number of goroutines that will increment the counter
)

func main() {
	wg := sync.WaitGroup{} // Initialize a WaitGroup to wait for all goroutines to finish

	// Launch numIncrements goroutines to increment the counter
	for i := 0; i < int(numIncrements); i++ {
		wg.Add(1)             // Add a goroutine to the WaitGroup
		go makeIncrement(&wg) // Start a new goroutine to increment the counter
	}

	wg.Wait() // Wait for all goroutines to complete

	fmt.Println("Final counter value:", counter) // Print the final value of the counter
}

// makeIncrement increments the shared counter atomically
func makeIncrement(wg *sync.WaitGroup) {
	defer wg.Done()              // Signal that this goroutine is done when the function returns
	atomic.AddInt64(&counter, 1) // Atomically add 1 to the counter to prevent race conditions
}
