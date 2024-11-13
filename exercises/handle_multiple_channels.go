package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels for even and odd numbers
	evenCh := make(chan int)
	oddCh := make(chan int)

	// Start goroutines to send even and odd numbers
	go sendEvenNumbers(evenCh)
	go sendOddNumbers(oddCh)

	// Flags to track if channels are closed
	evenClosed := false
	oddClosed := false

	// Use select to receive values from both channels concurrently
	for {
		// Exit the loop if both channels are closed
		if evenClosed && oddClosed {
			break
		}

		select {
		case evenNum, ok := <-evenCh:
			if !ok {
				evenClosed = true
				evenCh = nil // Disable this case
				continue
			}
			fmt.Println("Received from even channel:", evenNum)

		case oddNum, ok := <-oddCh:
			if !ok {
				oddClosed = true
				oddCh = nil // Disable this case
				continue
			}
			fmt.Println("Received from odd channel:", oddNum)
		}
	}

	fmt.Println("All numbers received!")
}

func sendEvenNumbers(ch chan<- int) {

	for i := 0; i < 10; i += 2 {
		time.Sleep(1 * time.Second) // Simulate task delay
		ch <- i
	}
	close(ch)
}

func sendOddNumbers(ch chan<- int) {

	for i := 1; i < 10; i += 2 {
		time.Sleep(1 * time.Second) // Simulate task delay
		ch <- i
	}
	close(ch)
}
