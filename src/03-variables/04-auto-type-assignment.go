package main

import "fmt"

func main() {
// START OMIT
First := "test"
Second := false
Third, Fourth := -2, 1.5
// END OMIT

	fmt.Printf("%-6s (%-7T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-7T): %-5v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-7T): %-5v\n", "Third", Third, Third)
	fmt.Printf("%-6s (%-7T): %-5v", "Fourth", Fourth, Fourth)
}
