package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create three WaitGroups to synchronize the goroutines
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	var wg3 sync.WaitGroup

	// Add a count of 1 to each WaitGroup
	wg.Add(1)
	wg2.Add(1)
	wg3.Add(1)

	// Start the first goroutine and wait for it to finish
	go firstPrint(&wg)
	wg.Wait()

	// Start the second goroutine and wait for it to finish
	go secondPrint(&wg2)
	wg2.Wait()

	// Start the third goroutine and wait for it to finish
	go thirdPrint(&wg3)
	wg3.Wait()

	// Print a message indicating that the main function has finished
	fmt.Println("Main function finished!")
}

// firstPrint prints "First!" and then signals that it is done
func firstPrint(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the function returns
	fmt.Println("First!")
}

// secondPrint prints "Second!" and then signals that it is done
func secondPrint(wg2 *sync.WaitGroup) {
	defer wg2.Done() // Decrement the WaitGroup counter when the function returns
	fmt.Println("Second!")
}

// thirdPrint prints "Third!" and then signals that it is done
func thirdPrint(wg3 *sync.WaitGroup) {
	defer wg3.Done() // Decrement the WaitGroup counter when the function returns
	fmt.Println("Third!")
}
