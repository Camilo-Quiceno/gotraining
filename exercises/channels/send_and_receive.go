package main

import (
	"fmt"
)

func main() {
	// Create a new channel of type string
	ch := make(chan string)

	// Start the senderMessage goroutine and pass the channel to it
	go senderMessage(ch)

	// Receive the message from the channel and assign it to the variable 'message'
	message := <-ch

	// Print the received message
	fmt.Println(message)

	// Indicate that the main goroutine has received the message
	fmt.Println("Main goroutine received the message!")
}

// senderMessage sends a message through the provided channel
func senderMessage(ch chan string) {
	defer close(ch)               // Ensure the channel is closed once the function completes
	ch <- "Hello from goroutine!" // Send the message to the channel
}
