package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {

	in := make(chan string)
	out := make(chan map[string]int)
	var wg sync.WaitGroup

	sentences := []string{
		"Concurrency in Go is powerful",
		"Goroutines simplify concurrent programming",
		"Channels enable communication between goroutines",
		"Synchronization is crucial for avoiding race conditions",
		"Fan-out pattern distributes tasks to multiple workers",
	}

	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go frequencyCounterWorker(in, out, &wg)
	}

	// Send sentences to the input channel
	go func() {
		for _, sentence := range sentences {
			in <- sentence
		}
		close(in) // Close the input channel after sending all sentences
	}()

	// Close the output channel after all workers have finished
	go func() {
		wg.Wait()
		close(out)
	}()

	var arrayOfMaps []map[string]int

	// Collect results from the output channel
	for wordCount := range out {
		arrayOfMaps = append(arrayOfMaps, wordCount)
	}

	// Combine all word count maps into a single map
	result := combineMaps(arrayOfMaps)

	// Print the word frequencies
	for key, value := range result {
		fmt.Printf("%s: %d\n", key, value)
	}

}

// frequencyCounterWorker processes sentences from the input channel,
// counts word frequencies, and sends the result to the output channel.
func frequencyCounterWorker(in <-chan string, out chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for sentence := range in {
		words := strings.Fields(sentence) // Split the sentence into words
		wordCount := make(map[string]int) // Create a map to hold word counts
		for _, word := range words {
			wordCount[word]++ // Increment the count for each word
		}
		out <- wordCount // Send the word count map to the output channel
	}
}

// combineMaps merges multiple word count maps into a single map
// by summing the counts of identical words.
func combineMaps(maps []map[string]int) map[string]int {
	// Create a map to store the combined results
	result := make(map[string]int)

	// Iterate over each map in the slice
	for _, m := range maps {
		// Iterate over each key-value pair in the map
		for key, value := range m {
			// Add the count to the corresponding word in the result map
			result[key] += value
		}
	}

	return result
}
