package main

import (
	"fmt"  // Import fmt package for formatted I/O operations
	"sync" // Import sync package for synchronization primitives like WaitGroup and Once
)

var (
	dbConnection string    // Shared variable to hold the database connection status
	cacheInit    string    // Shared variable to hold the cache initialization status
	loggerInit   string    // Shared variable to hold the logger initialization status
	onceDB       sync.Once // Ensures that the initializeDB function is executed only once
	onceCache    sync.Once // Ensures that the initializeCache function is executed only once
	onceLogger   sync.Once // Ensures that the initializeLogger function is executed only once
)

func main() {
	wg := sync.WaitGroup{} // Initialize a WaitGroup to wait for all goroutines to finish

	// Launch 10 goroutines to access the database
	for i := 0; i < 10; i++ {
		wg.Add(1)              // Increment the WaitGroup counter for each goroutine
		go accessDatabase(&wg) // Start a goroutine to access the database
	}

	// Launch 5 goroutines to access the cache
	for i := 0; i < 5; i++ {
		wg.Add(1)           // Increment the WaitGroup counter for each goroutine
		go accessCache(&wg) // Start a goroutine to access the cache
	}

	// Launch 5 goroutines to access the logger
	for i := 0; i < 5; i++ {
		wg.Add(1)            // Increment the WaitGroup counter for each goroutine
		go accessLogger(&wg) // Start a goroutine to access the logger
	}

	wg.Wait() // Wait for all goroutines to complete
}

// initializeDB sets up the database connection
func initializeDB() {
	dbConnection = "Database Connected!" // Update the dbConnection status
	fmt.Println("Database initialized")  // Print initialization message
}

// initializeCache sets up the cache system
func initializeCache() {
	cacheInit = "Cache Initialized!" // Update the cacheInit status
	fmt.Println("Cache initialized") // Print initialization message
}

// initializeLogger sets up the logger system
func initializeLogger() {
	loggerInit = "Logger Initialized!" // Update the loggerInit status
	fmt.Println("Logger initialized")  // Print initialization message
}

// accessDatabase ensures the database is initialized and then accesses it
func accessDatabase(wg *sync.WaitGroup) {
	defer wg.Done()                         // Signal that this goroutine is done when the function returns
	onceDB.Do(initializeDB)                 // Ensure initializeDB is called only once
	fmt.Println("Accessing:", dbConnection) // Print the current state of the database connection
}

// accessCache ensures the database is initialized before initializing the cache, then accesses it
func accessCache(wg *sync.WaitGroup) {
	accessDatabase(wg)                   // Ensure the database is initialized first
	onceCache.Do(initializeCache)        // Ensure initializeCache is called only once
	fmt.Println("Accessing:", cacheInit) // Print the current state of the cache
}

// accessLogger ensures the logger is initialized and then accesses it
func accessLogger(wg *sync.WaitGroup) {
	defer wg.Done()                       // Signal that this goroutine is done when the function returns
	onceLogger.Do(initializeLogger)       // Ensure initializeLogger is called only once
	fmt.Println("Accessing:", loggerInit) // Print the current state of the logger
}
