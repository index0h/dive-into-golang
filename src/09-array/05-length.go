package main

import "fmt"

func main() {
// START OMIT
First := [...]int8{1, 2, 3, 4}
Length := len(First)
// END OMIT

	fmt.Printf("%-6s (%-7T): %v\n", "First", First, First)
	fmt.Printf("%-6s (%-7T): %v", "Length", Length, Length)
}
