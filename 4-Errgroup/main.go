package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// Create a new context and an errgroup.Group
	ctx := context.Background()
	group, ctx := errgroup.WithContext(ctx)

	// Example goroutine 1
	group.Go(func() error {
		time.Sleep(2 * time.Second) // Simulate some work
		fmt.Println("Goroutine 1 completed")
		return nil
	})

	// Example goroutine 2 with an error
	group.Go(func() error {
		time.Sleep(3 * time.Second) // Simulate some work
		fmt.Println("Goroutine 2 completed")
		return fmt.Errorf("goroutine 2 encountered an error")
	})

	// Example goroutine 3
	group.Go(func() error {
		time.Sleep(1 * time.Second) // Simulate some work
		fmt.Println("Goroutine 3 completed")
		return nil
	})

	// Wait for all goroutines to complete and check for errors
	if err := group.Wait(); err != nil {
		fmt.Printf("One or more goroutines encountered errors: %v\n", err)
	} else {
		fmt.Println("All goroutines completed successfully")
	}
}
