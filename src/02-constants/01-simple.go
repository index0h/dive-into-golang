package main

import "fmt"
// START OMIT
const First = "first"

const (
	Second = 2
	Third = "third"
)
// END OMIT
func main() {
	fmt.Printf("%-6s (%-6T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-6T): %-5v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-6T): %-5v", "Third", Third, Third)
}

