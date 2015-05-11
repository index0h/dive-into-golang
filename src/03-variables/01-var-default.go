package main

import "fmt"

func main() {
// START OMIT
var First string
var Second, Third bool
var (
	Fourth int
	Fifth float32
)
// END OMIT

	fmt.Printf("%-6s (%-7T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-7T): %-5v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-7T): %-5v\n", "Third", Third, Third)
	fmt.Printf("%-6s (%-7T): %-5v\n", "Fourth", Fourth, Fourth)
	fmt.Printf("%-6s (%-7T): %-5v", "Fifth", Fifth, Fifth)
}
