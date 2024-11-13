package main

import (
	"fmt"  // Import fmt package for formatted I/O
	"sync" // Import sync package for synchronization primitives
)

var mutex sync.Mutex // Mutex to protect access to the counter
var counter int      // Shared counter variable

func main() {
	wg := sync.WaitGroup{} // Initialize a WaitGroup to wait for all goroutines to finish
	wg.Add(5)              // Add five goroutines to the WaitGroup

	// Start five goroutines to increment the counter
	go incrementCounter1(&wg)
	go incrementCounter2(&wg)
	go incrementCounter3(&wg)
	go incrementCounter4(&wg)
	go incrementCounter5(&wg)

	wg.Wait() // Wait for all goroutines to complete

	fmt.Println("Final counter value:", counter) // Print the final value of the counter
}

// incrementCounter1 increments the counter by 1000
func incrementCounter1(wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done when the function returns
	mutex.Lock()    // Acquire the mutex before accessing the counter
	counter += 1000 // Increment the counter
	mutex.Unlock()  // Release the mutex after updating the counter
}

// incrementCounter2 increments the counter by 1000
func incrementCounter2(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter += 1000
	mutex.Unlock()
}

// incrementCounter3 increments the counter by 1000
func incrementCounter3(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter += 1000
	mutex.Unlock()
}

// incrementCounter4 increments the counter by 1000
func incrementCounter4(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter += 1000
	mutex.Unlock()
}

// incrementCounter5 increments the counter by 1000
func incrementCounter5(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter += 1000
	mutex.Unlock()
}
