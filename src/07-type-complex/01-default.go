package main

import "fmt"

func main() {
// START OMIT
var Value complex64
Type := 1 + 2i
// END OMIT

	fmt.Printf("%-5s (%-10T): %v\n", "Value", Value, Value)
	fmt.Printf("%-5s (%-10T): %v", "Type", Type, Type)
}
