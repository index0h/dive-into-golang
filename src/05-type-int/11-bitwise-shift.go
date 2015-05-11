package main

import "fmt"

func main() {
	var x, left, right uint8

// START OMIT
x = 9
left = x << 2
right = x >> 2
// END OMIT

	fmt.Printf("%-5v = %.8b = %-2d\n", "x", x, x)
	fmt.Printf("%-5v = %.8b = %-2d\n", "left", left, left)
	fmt.Printf("%-5v = %.8b = %-2d", "right", right, right)
}
