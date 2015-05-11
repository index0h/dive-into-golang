package main

import "fmt"

func main() {
// START OMIT
First := []byte{1, 2, 3}
Second := First

Second[0] *= 10
Second[1] *= 10
Second[2] *= 10
// END OMIT

	fmt.Printf("%-6s (%T): %v\n", "First", First, First)
	fmt.Printf("%-6s (%T): %v", "Second", Second, Second)
}
