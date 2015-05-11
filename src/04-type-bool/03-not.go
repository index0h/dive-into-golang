package main

import "fmt"

func main() {
// START OMIT
First := !true
Second := !false
// END OMIT

	fmt.Printf("%-6s (%-4T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-4T): %-5v", "Second", Second, Second)
}
