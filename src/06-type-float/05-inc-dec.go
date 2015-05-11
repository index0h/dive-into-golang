package main

import "fmt"

func main() {
	var Increment, Decrement float32

// START OMIT
Increment = 5.1
Increment++

Decrement = 5.2
Decrement--
// END OMIT

	fmt.Printf("%-8s (%-3T): %-5v\n", "Increment", Increment, Increment)
	fmt.Printf("%-8s (%-3T): %-5v", "Decrement", Decrement, Decrement)
}
