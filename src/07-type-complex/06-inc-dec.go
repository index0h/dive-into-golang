package main

import "fmt"

func main() {
	var Increment, Decrement complex64

// START OMIT
Increment = 5.1 + 2i
Increment++

Decrement = 5.2 + 2i
Decrement--
// END OMIT

	fmt.Printf("%-8s (%-3T): %v\n", "Increment", Increment, Increment)
	fmt.Printf("%-8s (%-3T): %v", "Decrement", Decrement, Decrement)
}
