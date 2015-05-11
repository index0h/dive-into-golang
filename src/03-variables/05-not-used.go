package main

import "fmt"

func main() {
// START OMIT
_, Used := "not used data", "used data"
// NotUsed, Used := "not used data", "used data" // Error
// END OMIT

	fmt.Printf("%-4s (%-6T): %-5v", "Used", Used, Used)
}
