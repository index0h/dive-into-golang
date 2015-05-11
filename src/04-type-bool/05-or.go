package main

import "fmt"

func main() {
// START OMIT
First := true || true
Second := true || false
Third := false || false
// END OMIT

	fmt.Printf("%-6s (%-4T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-4T): %-5v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-4T): %-5v", "Third", Third, Third)
}
