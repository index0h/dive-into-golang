package main

import "fmt"

func main() {
// START OMIT
var First string
Second := "line1\nline2"
Third := `line1\nline2`
// END OMIT

	fmt.Printf("%-6s (%T): '%v'\n\n", "First", First, First)
	fmt.Printf("%-6s (%T):\nvalue: '%v'\n\n", "Second", Second, Second)
	fmt.Printf("%-6s (%T):\nvalue: '%v'", "Third", Third, Third)
}
