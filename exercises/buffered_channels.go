package main

import "fmt"

func main() {
	// Create a buffered channel with a capacity of 3 integers
	c := make(chan int, 3)

	// Start the createMessages goroutine and pass the channel to it
	go createMessages(c)

	fmt.Println("--- Receiving from channel ---")
	// Receive messages from the channel until it is closed
	for msg := range c {
		fmt.Println("Received", msg) // Print each received message
	}

}

// createMessages sends integers to the provided channel
func createMessages(c chan int) {
	defer close(c) // Ensure the channel is closed after sending all messages

	// Send integers 1 to 3 to the channel
	for i := 1; i < 4; i++ {
		fmt.Println("Sending", i, " to channel...") // Print the integer being sent
		c <- i                                      // Send the integer to the channel
		fmt.Println("Sent", i, " to channel...")
	}
}
