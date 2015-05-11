package main

import "fmt"

func main() {
// START OMIT
var Value float32
Type := 13.2
// END OMIT

	fmt.Printf("%-5s (%T): %-5v\n", "Value", Value, Value)
	fmt.Printf("%-5s (%T): %-5v", "Type", Type, Type)
}
