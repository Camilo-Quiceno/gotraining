package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a buffered channel with a capacity of 3 strings
	c := make(chan string, 3)

	// Create a WaitGroup to wait for both goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2) // We have two goroutines to wait for

	// Start the sendMessages goroutine and pass the channel and WaitGroup
	go sendMessages(c, &wg)

	// Start the receiveMessages goroutine and pass the channel and WaitGroup
	go receiveMessages(c, &wg)

	// Wait for both goroutines to finish
	wg.Wait()

	fmt.Println("All messages received!")
}

// sendMessages sends messages to the channel and signals the WaitGroup when done
func sendMessages(c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	defer close(c)  // Close the channel after sending all messages

	fmt.Println("Sending: Hello")
	c <- "Hello"

	fmt.Println("Sending: World")
	c <- "World"

	fmt.Println("Sending: Golang")
	c <- "Golang"
}

// receiveMessages receives messages from the channel and signals the WaitGroup when done
func receiveMessages(c <-chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done

	for msg := range c {
		fmt.Println("Received", msg)
	}
}
