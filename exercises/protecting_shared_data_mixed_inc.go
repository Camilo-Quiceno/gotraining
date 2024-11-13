package main

import (
	"fmt"
	"sync"
)

var (
	mutex                 sync.Mutex // Mutex to protect access to the counter
	counter               = 0        // Shared counter variable
	numberIncreasers      = 5        // Number of goroutines that will increment the counter
	numberDecreasers      = 3        // Number of goroutines that will decrement the counter
	operationPerGoroutine = 1000     // Number of operations each goroutine will perform
)

func main() {
	wg := sync.WaitGroup{} // Initialize a WaitGroup to wait for all goroutines to finish

	// Start incrementing goroutines
	for i := 0; i < numberIncreasers; i++ {
		wg.Add(1) // Add a goroutine to the WaitGroup
		go incrementCounter(&wg)
	}

	// Start decrementing goroutines
	for i := 0; i < numberDecreasers; i++ {
		wg.Add(1) // Add a goroutine to the WaitGroup
		go decrementCounter(&wg)
	}

	wg.Wait() // Wait for all goroutines to complete

	fmt.Println("Final counter value:", counter) // Print the final value of the counter
}

// incrementCounter increments the shared counter by a predefined amount
func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()                  // Signal that the goroutine is done when the function returns
	mutex.Lock()                     // Acquire the mutex before accessing the counter
	counter += operationPerGoroutine // Increment the counter
	mutex.Unlock()                   // Release the mutex after updating the counter
}

// decrementCounter decrements the shared counter by a predefined amount
func decrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()                  // Signal that the goroutine is done when the function returns
	mutex.Lock()                     // Acquire the mutex before accessing the counter
	counter -= operationPerGoroutine // Decrement the counter
	mutex.Unlock()                   // Release the mutex after updating the counter
}
