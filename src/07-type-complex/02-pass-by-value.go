package main

import "fmt"

func main() {
// START OMIT
First := complex64(2.1 + 2i)
Second := First
Second = complex64(5.1 + 1i)
// END OMIT

	fmt.Printf("%-6s (%T): %v\n", "First", First, First)
	fmt.Printf("%-6s (%T): %v", "Second", Second, Second)
}
