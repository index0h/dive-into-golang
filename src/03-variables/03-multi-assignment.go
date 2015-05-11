package main

import "fmt"

func main() {
// START OMIT
var (
	First string
	Second int
	Third float32
)
First, Second, Third = "test", -12, 5.4
// END OMIT

	fmt.Printf("%-6s (%-7T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-7T): %-5v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-7T): %-5v", "Third", Third, Third)
}
