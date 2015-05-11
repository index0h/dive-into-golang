package main

import "fmt"

func main() {
	var Increment, Decrement int

// START OMIT
Increment = 5
Increment++

Decrement = 5
Decrement--
// END OMIT

	fmt.Printf("%-8s (%-3T): %-5v\n", "Increment", Increment, Increment)
	fmt.Printf("%-8s (%-3T): %-5v", "Decrement", Decrement, Decrement)
}
