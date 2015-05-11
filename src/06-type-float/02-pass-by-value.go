package main

import "fmt"

func main() {
// START OMIT
First := float32(2.1)
Second := First
Second = float32(5.1)
// END OMIT

	fmt.Printf("%-6s (%T): %v\n", "First", First, First)
	fmt.Printf("%-6s (%T): %v", "Second", Second, Second)
}
