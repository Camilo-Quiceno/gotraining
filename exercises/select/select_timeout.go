package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	go sendFunction(c)

	for {
		select {
		case ms := <-c:
			fmt.Println(ms)

		case <-time.After(1 * time.Second):
			fmt.Println("Timeout! No message received.!")
			return // Exit the loop on timeout
		}
	}

}

func sendFunction(c chan<- string) {
	fmt.Println("Sending message to channel...")
	time.Sleep(3 * time.Second) // Simulate task delay
	c <- "Task complete!"
	fmt.Println("Message sent to channel.")
}
