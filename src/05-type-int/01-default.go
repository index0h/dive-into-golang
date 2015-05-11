package main

import "fmt"

func main() {
// START OMIT
var Value int
Type := 15
// END OMIT

	fmt.Printf("%-5s (%-3T): %-5v\n", "Value", Value, Value)
	fmt.Printf("%-5s (%-3T): %-5v", "Type", Type, Type)
}
