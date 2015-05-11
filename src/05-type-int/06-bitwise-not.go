package main

import "fmt"

func main() {
	var x, not uint8

// START OMIT
x = 9
not = ^x
// END OMIT

	fmt.Printf("%-3v = %.8b = %-3d\n", "x", x, x)
	fmt.Printf("%-3v = %.8b = %-3d", "not", not, not)
}
