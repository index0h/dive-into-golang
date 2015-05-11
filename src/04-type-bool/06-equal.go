package main

import "fmt"

func main() {
// START OMIT
First := 1 == 2
Second := 1 != 2
// END OMIT

	fmt.Printf("%-6s (%-4T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-4T): %-5v", "Second", Second, Second)
}
