package main

import "fmt"

func main() {
// START OMIT
var First string = "test"
var (
	Second int = 15
	Third float32 = -10.25
)
// END OMIT

	fmt.Printf("%-6s (%-7T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-7T): %-5v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-7T): %-5v", "Third", Third, Third)
}
