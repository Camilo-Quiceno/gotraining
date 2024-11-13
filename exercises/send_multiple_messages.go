package main

import (
	"fmt"
)

func main() {
	// Create a new channel of type string
	ch := make(chan string)

	// Start the senderMessage goroutine and pass the channel to it
	go senderMessage(ch)

	// Receive messages from the channel until it is closed
	for message := range ch {
		fmt.Println(message) // Print each received message
	}

	// Print a message indicating that all messages have been received
	fmt.Println("All messages received!")
}

// senderMessage sends multiple messages through the provided channel
func senderMessage(ch chan string) {
	defer close(ch) // Ensure the channel is closed once all messages are sent

	// Loop to send messages from "Message 1" to "Message 5"
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Message %d", i) // Send the formatted message to the channel
	}
}
