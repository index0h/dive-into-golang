package main

import "fmt"

func main() {
// START OMIT
First := "abc"
Second := First

Second = "DEF"
// END OMIT

	fmt.Printf("%-6s (%T): %v\n", "First", First, First)
	fmt.Printf("%-6s (%T): %v", "Second", Second, Second)
}
