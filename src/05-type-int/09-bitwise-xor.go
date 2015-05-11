package main

import "fmt"

func main() {
	var x, y, xor uint8

// START OMIT
x = 9
y = 28
xor = x ^ y
// END OMIT

	fmt.Printf("%-3v = %.8b = %-3d\n", "x", x, x)
	fmt.Printf("%-3v = %.8b = %-3d\n", "y", y, y)
	fmt.Printf("%-3v = %.8b = %-3d", "xor", xor, xor)
}
