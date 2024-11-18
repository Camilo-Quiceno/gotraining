package main

import "fmt"

// main is the entry point of the program.
func main() {
	num := 10                    // Define the number of integers to generate.
	nums := numberGenerator(num) // Create a channel that will receive generated numbers.
	squares := squating(nums)    // Create a channel that will receive squared numbers.
	printNumber(squares)         // Start printing the squared numbers.
}

// numberGenerator generates integers from 1 to n and sends them to the returned channel.
func numberGenerator(n int) <-chan int {
	out := make(chan int) // Create a new channel to send integers.

	// Start a goroutine to generate numbers.
	go func() {
		for i := 0; i < n; i++ {
			out <- i + 1 // Send the number (i + 1) to the channel.
		}
		close(out) // Close the channel after sending all numbers.
	}()

	return out // Return the channel to the caller.
}

// squating receives integers from the input channel, squares them,
// and sends the results to the returned channel.
func squating(in <-chan int) <-chan int {
	out := make(chan int) // Create a new channel to send squared numbers.

	// Start a goroutine to process and square the numbers.
	go func() {
		for n := range in { // Receive numbers from the input channel.
			out <- n * n // Send the squared number to the output channel.
		}
		close(out) // Close the channel after processing all numbers.
	}()

	return out // Return the channel to the caller.
}

// printNumber receives squared numbers from the input channel and prints them.
func printNumber(in <-chan int) {
	for n := range in { // Receive squared numbers from the channel.
		fmt.Println("Squared Number:", n) // Print each squared number.
	}
}
