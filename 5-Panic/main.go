package main

import (
	"fmt"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func process() {
	defer recoverFromPanic() // Defer the recover function to handle panics

	fmt.Println("Processing step 1")
	// Simulate a panic condition
	panic("Something went wrong!")

	fmt.Println("Processing step 2") // This line won't be executed due to the panic
}

func main() {
	fmt.Println("Start of main function")

	process() // Call the process function that may panic

	fmt.Println("End of main function") // This line won't be executed due to the panic
}
