package main

import "fmt"

func main() {
// START OMIT
var First [2]bool
Second := [...]int{-1, -2, -3}
Third := [3]uint{1, 2, 3}
// END OMIT

	fmt.Printf("%-6s (%-7T): %v\n", "First", First, First)
	fmt.Printf("%-6s (%-7T): %v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-7T): %v", "Third", Third, Third)
}
