package main

import "fmt"

// main is the entry point of the program.
func main() {
	num := 100                   // Define the upper limit for prime number generation.
	nums := generateNumbers(num) // Initialize the number generator and receive a channel of numbers.
	primes := primeCheck(nums)   // Initialize the prime checker and receive a channel of prime numbers.
	primeCollector(primes)       // Collect and print the prime numbers from the primes channel.
}

// generateNumbers creates a channel and sends integers from 1 up to n.
func generateNumbers(n int) <-chan int {
	out := make(chan int) // Create a new unbuffered channel of integers.

	// Start a goroutine to generate numbers concurrently.
	go func() {
		for i := 0; i < n; i++ { // Loop from 0 to n-1.
			out <- i + 1 // Send the number (i + 1) to the out channel.
		}
		close(out) // Close the channel to signal that no more numbers will be sent.
	}()

	return out // Return the channel to the caller.
}

// primeCheck receives integers from the input channel, checks for primality,
// and sends prime numbers to the output channel.
func primeCheck(in <-chan int) <-chan int {
	out := make(chan int) // Create a new unbuffered channel for prime numbers.

	// Start a goroutine to process numbers concurrently.
	go func() {
		for n := range in { // Receive numbers from the input channel until it's closed.
			if isPrime(n) { // Check if the current number is prime.
				out <- n // Send the prime number to the out channel.
			}
		}
		close(out) // Close the out channel to signal that no more primes will be sent.
	}()

	return out // Return the channel to the caller.
}

// primeCollector receives prime numbers from the input channel,
// stores them in a slice, and prints the list of primes.
func primeCollector(in <-chan int) {
	primes := []int{} // Initialize an empty slice to store prime numbers.

	for n := range in { // Receive prime numbers from the input channel until it's closed.
		primes = append(primes, n) // Append each received prime number to the primes slice.
	}

	// Print the list of collected prime numbers.
	fmt.Println("Prime Numbers up to 100:")
	for i, p := range primes { // Iterate over the primes slice with index and value.
		if i > 0 {
			fmt.Print(", ") // Print a comma before all primes except the first one for formatting.
		}
		fmt.Print(p) // Print the prime number.
	}
	fmt.Println() // Print a newline at the end for clean output.
}

// isPrime determines whether a given integer n is a prime number.
// It returns true if n is prime, and false otherwise.
func isPrime(n int) bool {
	if n <= 1 {
		return false // Numbers less than or equal to 1 are not prime.
	}

	for i := 2; i*i <= n; i++ { // Only need to check divisors up to the square root of n.
		if n%i == 0 {
			return false // If n is divisible by i, it's not prime.
		}
	}

	return true // If no divisors found, n is prime.
}
