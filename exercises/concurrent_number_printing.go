package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a WaitGroup to synchronize the goroutines
	var wg sync.WaitGroup

	// Add a count of 5 to the WaitGroup, one for each goroutine
	wg.Add(5)

	// Start five goroutines to print numbers from 1 to 5
	go printOne(&wg)
	go printTwo(&wg)
	go printThree(&wg)
	go printFour(&wg)
	go printFive(&wg)

	// Wait for all goroutines to finish
	wg.Wait()

	// Print a message indicating that all goroutines have finished
	fmt.Println("All goroutines finished!")
}

// printOne prints "1" and then signals that it is done
func printOne(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the function returns
	fmt.Println("1")
}

// printTwo prints "2" and then signals that it is done
func printTwo(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the function returns
	fmt.Println("2")
}

// printThree prints "3" and then signals that it is done
func printThree(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the function returns
	fmt.Println("3")
}

// printFour prints "4" and then signals that it is done
func printFour(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the function returns
	fmt.Println("4")
}

// printFive prints "5" and then signals that it is done
func printFive(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the function returns
	fmt.Println("5")
}
