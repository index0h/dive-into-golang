package main

import "fmt"

func main() {
	var x, y, mod uint8

// START OMIT
x = 9
y = 28
mod = x % y
// END OMIT

	fmt.Printf("%-5v = %.8b = %-2d\n", "x", x, x)
	fmt.Printf("%-5v = %.8b = %-2d\n", "y", y, y)
	fmt.Printf("%-5v = %.8b = %-2d", "mod", mod, mod)
}
