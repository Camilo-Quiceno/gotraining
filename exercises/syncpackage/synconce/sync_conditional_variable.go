package main

import (
	"fmt"  // Import fmt package for formatted I/O operations
	"sync" // Import sync package for synchronization primitives like WaitGroup and Once
)

var (
	data string    // Shared data variable that will be initialized once
	once sync.Once // Ensures that a function is only executed once
)

func main() {
	wg := sync.WaitGroup{} // Initialize a WaitGroup to wait for all goroutines to finish

	// Launch 5 goroutines to print data
	for i := 0; i < 5; i++ {
		wg.Add(1)         // Increment the WaitGroup counter
		go printData(&wg) // Start a new goroutine to execute printData
	}

	wg.Wait() // Block until all goroutines have called wg.Done()
}

// initializeData initializes the shared data variable
func initializeData() {
	data = "Initialized!"           // Set the shared data
	fmt.Println("Data initialized") // Print a message indicating initialization
}

// printData ensures data is initialized and then prints it
func printData(wg *sync.WaitGroup) {
	defer wg.Done()            // Decrement the WaitGroup counter when the function returns
	once.Do(initializeData)    // Ensure initializeData is called only once
	fmt.Println("Data:", data) // Print the current value of data
}
