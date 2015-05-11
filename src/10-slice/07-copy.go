package main

import "fmt"

func main() {
// START OMIT
First := []int{1, 2, 3}
Second := make([]int, 2, 3)

Count := copy(Second, First)

First[0] *= 5
First[1] *= 5
First[2] *= 5
// END OMIT

	fmt.Printf("%-6s (%T): %v\n\n", "Count", Count, Count)
	fmt.Printf("%-6s (%T): len: %v cap: %v\nvalue: %v\n\n", "First", First, len(First), cap(First), First)
	fmt.Printf("%-6s (%T): len: %v cap: %v\nvalue: %v", "Second", Second, len(Second), cap(Second), Second)
}
