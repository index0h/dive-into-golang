package main

import "fmt"

func main() {
// START OMIT
First := int(2)
Second := First
Second = int(5)
// END OMIT

	fmt.Printf("%-6s (%T): %v\n", "First", First, First)
	fmt.Printf("%-6s (%T): %v", "Second", Second, Second)
}
