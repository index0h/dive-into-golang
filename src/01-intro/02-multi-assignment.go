package main

import "fmt"

func main() {
	var First string
	var Second bool

// START OMIT
First, Second = "some string", true
// END OMIT

	fmt.Printf("%-6s (%-6T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-6T): %-5v\n", "Second", Second, Second)
}
