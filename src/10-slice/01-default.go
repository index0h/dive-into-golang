package main

import "fmt"

func main() {
// START OMIT
First := make([]byte, 5, 15)
Second := make([]byte, 3)
var Third []bool
Fourth := []int{1, 2, 3, 4}
// END OMIT

	fmt.Printf("%-6s (%-7T): %v\n", "First", First, First)
	fmt.Printf("%-6s (%-7T): %v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-7T): %v\n", "Third", Third, Third)
	fmt.Printf("%-6s (%-7T): %v", "Fourth", Fourth, Fourth)
}
