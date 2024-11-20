package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	numWorkers = 3  // Number of producer goroutines
	numNumbers = 10 // Number of numbers each producer will generate
)

func main() {
	var wg sync.WaitGroup // WaitGroup to synchronize producer goroutines

	var channels []<-chan int
	// Start multiple producer goroutines and collect their output channels
	for i := 0; i < numWorkers; i++ {
		channels = append(channels, producer(numNumbers, &wg))
	}

	out := fanIn(channels...) // Merge the output from all producers into a single channel
	total := 0                // Variable to hold the total sum of all numbers

	// Read from the merged output channel and accumulate the total
	for n := range out {
		total += n
	}
	fmt.Println("Total Sum:", total) // Print the total sum

	wg.Wait() // Wait for all producer goroutines to finish
}

// producer generates numbers from 0 to n, sending each number to a channel.
// It simulates work by sleeping for a random duration between sends.
func producer(n int, wg *sync.WaitGroup) <-chan int {
	out := make(chan int) // Create an unbuffered channel for sending numbers
	wg.Add(1)             // Increment the WaitGroup counter to indicate a new goroutine

	go func() {
		defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
		for i := 0; i < n+1; i++ {
			// Generate a random sleep duration between 0 to 999 milliseconds
			randomDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
			time.Sleep(randomDuration) // Simulate work by sleeping
			out <- i                   // Send the current number to the output channel
		}
		close(out) // Close the channel after sending all numbers to signal completion
	}()
	return out // Return the channel to the caller so it can receive the numbers
}

// fanIn merges multiple input channels into a single output channel.
// It starts a goroutine for each input channel to read and send its data to the output channel.
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int) // Create an unbuffered channel for merged output
	var wg sync.WaitGroup // WaitGroup to synchronize fan-in goroutines

	// Start a goroutine for each input channel
	for _, c := range channels {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go func(c <-chan int) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			total := 0      // Variable to accumulate the sum from this channel
			for n := range c {
				total += n // Add each number received to the total
			}
			out <- total // Send the total sum from this channel to the merged output channel
		}(c)
	}

	// Start a goroutine to close the output channel once all fan-in goroutines have finished
	go func() {
		wg.Wait()  // Wait for all fan-in goroutines to finish
		close(out) // Close the merged output channel to signal that no more data will be sent
	}()
	return out // Return the merged output channel to the caller
}
