package main

import "fmt"
// START OMIT
const (
	First = 1
	Second
	Third = "third"
	Fourth
)
// END OMIT

func main() {
	fmt.Printf("%-6s (%-6T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-6T): %-5v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-6T): %-5v\n", "Third", Third, Third)
	fmt.Printf("%-6s (%-6T): %-5v", "Fourth", Fourth, Fourth)
}
