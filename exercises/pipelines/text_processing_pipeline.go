package main

import "strings"

// main is the entry point of the program.
func main() {
	lines := []string{
		"INFO: Starting the application",
		"DEBUG: Initializing modules",
		"ERROR: Failed to connect to database",
		"INFO: Application running",
		"ERROR: Null pointer exception encountered",
		"DEBUG: Shutting down modules",
		"INFO: Application terminated",
	}

	lineCh := lineGenerator(lines)         // Start the line generator and receive a channel of lines.
	errorCh := filterLines(lineCh)         // Filter lines containing "ERROR" and receive a channel of filtered lines.
	upperCh := convertToUpperCase(errorCh) // Convert filtered lines to uppercase and receive a channel of transformed lines.
	printLines(upperCh)                    // Print the transformed lines from the final channel.
}

// lineGenerator sends each line from the provided slice to the output channel.
func lineGenerator(lines []string) <-chan string {
	out := make(chan string) // Create an unbuffered channel for sending lines.
	go func() {
		for _, line := range lines { // Iterate over each line in the slice.
			out <- line // Send the current line to the output channel.
		}
		close(out) // Close the channel after all lines have been sent.
	}()
	return out // Return the output channel to the caller.
}

// filterLines receives lines from the input channel, filters those containing "ERROR", and sends them to the output channel.
func filterLines(in <-chan string) <-chan string {
	out := make(chan string) // Create an unbuffered channel for sending filtered lines.
	go func() {
		for line := range in { // Receive lines from the input channel until it's closed.
			if strings.Contains(strings.ToUpper(line), "ERROR") { // Check if the line contains "ERROR" (case-insensitive).
				out <- line // Send the filtered line to the output channel.
			}
		}
		close(out) // Close the output channel after filtering is complete.
	}()
	return out // Return the filtered lines channel to the caller.
}

// convertToUpperCase receives lines from the input channel, converts them to uppercase, and sends them to the output channel.
func convertToUpperCase(in <-chan string) <-chan string {
	out := make(chan string) // Create an unbuffered channel for sending transformed lines.
	go func() {
		for line := range in { // Receive lines from the input channel until it's closed.
			out <- strings.ToUpper(line) // Convert the line to uppercase and send it to the output channel.
		}
		close(out) // Close the output channel after all lines have been transformed.
	}()
	return out // Return the transformed lines channel to the caller.
}

// printLines receives lines from the input channel and prints each line to the console.
func printLines(in <-chan string) {
	for line := range in { // Receive lines from the input channel until it's closed.
		println(line) // Print the current line to the console.
	}
}
