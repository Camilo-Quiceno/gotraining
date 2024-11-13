package main

import (
	"fmt"
	"time"
)

func main() {

	go firstPrint()
	go secondPrint()

	time.Sleep(1 * time.Second)
	fmt.Println("Finishing main function")
}

func firstPrint() {
	fmt.Println("Hello from goroutine 1!")
}

func secondPrint() {
	fmt.Println("Hello from goroutine 2!")
}
