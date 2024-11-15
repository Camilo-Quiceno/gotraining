package main

import (
	"fmt"       // Import fmt package for formatted I/O operations
	"math/rand" // Import math/rand package for generating random numbers
	"sync"      // Import sync package for synchronization primitives like WaitGroup and Once
	"time"      // Import time package for handling time-related functions
)

var (
	dbConnection string    // Shared variable to hold the database connection status
	once         sync.Once // Ensures that the initializeDB function is executed only once
)

func main() {
	wg := sync.WaitGroup{} // Initialize a WaitGroup to wait for all goroutines to finish

	for i := 0; i < 10; i++ { // Loop to launch 10 goroutines
		wg.Add(1)                                            // Increment the WaitGroup counter for each goroutine
		randomValue := rand.Intn(10) + 1                     // Generate a random integer between 1 and 10
		time.Sleep(time.Duration(randomValue) * time.Second) // Sleep for randomValue seconds before launching the goroutine
		go accessDatabase(&wg)                               // Launch a goroutine to access the database
	}
	wg.Wait() // Wait for all goroutines to complete
}

// initializeDB initializes the database connection
func initializeDB() {
	dbConnection = "Database Connected!" // Set the dbConnection string to indicate a successful connection
	fmt.Println("Database initialized")  // Print a message indicating that the database has been initialized
}

// accessDatabase ensures the database is initialized and then accesses it
func accessDatabase(wg *sync.WaitGroup) {
	defer wg.Done()                         // Signal that this goroutine is done when the function returns
	once.Do(initializeDB)                   // Ensure that initializeDB is called only once, even if multiple goroutines reach this point
	fmt.Println("Accessing:", dbConnection) // Print the current state of the database connection
}
