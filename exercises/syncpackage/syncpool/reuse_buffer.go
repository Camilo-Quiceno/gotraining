package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Buffer represents a reusable buffer with a slice of bytes.
type Buffer struct {
	Data []byte
}

func main() {
	wg := sync.WaitGroup{} // Initialize a WaitGroup to wait for all goroutines to finish.

	// Initialize a sync.Pool to manage reusable Buffer objects.
	buffer := sync.Pool{
		New: func() interface{} {
			// Define how to create a new Buffer if the pool is empty.
			return &Buffer{
				Data: []byte("Buffer data - Hello from worker"),
			}
		},
	}

	// Launch 2 worker goroutines.
	for i := 1; i < 3; i++ {
		wg.Add(1)                          // Increment the WaitGroup counter for each goroutine.
		go workerFunction(&wg, &buffer, i) // Start a worker goroutine with its ID.
	}

	wg.Wait() // Wait for all goroutines to complete.
}

// workerFunction retrieves a Buffer from the pool, modifies it, prints the result,
// and then returns the Buffer to the pool.
func workerFunction(wg *sync.WaitGroup, buffer *sync.Pool, num int) {
	defer wg.Done() // Signal that this goroutine is done when the function returns.

	obj := buffer.Get().(*Buffer) // Retrieve a Buffer from the pool and assert its type.

	ogData := obj.Data // Store the original data of the Buffer.

	// Append the worker's ID to the Buffer's data.
	obj.Data = []byte(string(obj.Data) + " " + strconv.Itoa(num))
	fmt.Println(string(obj.Data)) // Print the modified Buffer data.

	obj.Data = ogData // Reset the Buffer's data to its original state.

	buffer.Put(obj) // Return the Buffer to the pool for reuse.
}
