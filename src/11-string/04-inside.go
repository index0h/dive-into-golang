package main

import "fmt"

func main() {
// START OMIT
First := "DF"
Second := "БГ"

BytesFirst := []byte(First)
BytesSecond := []byte(Second)
// END OMIT

	fmt.Printf("%-6s (%T): len: %v\nvalue: %v\nbytes: %v\n\n", "First", First, len(First), First, BytesFirst)
	fmt.Printf("%-6s (%T): len: %v\nvalue: %v\nbytes: %v", "Second", Second, len(Second), Second, BytesSecond)
}
