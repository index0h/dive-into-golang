package main

import "fmt"

func main() {
// START OMIT
var Value bool
// END OMIT

	fmt.Printf("%s (%-4T): %-5v", "Value", Value, Value)
}
