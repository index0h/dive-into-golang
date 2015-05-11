package main

import "fmt"

func main() {
// START OMIT
First := []int{1, 2, 3}
Second := []int{4, 5, 6}

First = append(First, Second[:2]...)
Second = append(Second, 7, 8)
// END OMIT

	fmt.Printf("%-6s (%T): len: %v cap: %v\nvalue: %v\n\n", "First", First, len(First), cap(First), First)
	fmt.Printf("%-6s (%T): len: %v cap: %v\nvalue: %v", "Second", Second, len(Second), cap(Second), Second)
}
