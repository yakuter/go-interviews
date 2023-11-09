package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)

	// go func() {
	g1 := <-c1 // wait for value
	fmt.Println("get g1: ", g1)
	// }()

	fmt.Println("push c1: ")
	c1 <- 10 // send value and wait until it is received.
}
