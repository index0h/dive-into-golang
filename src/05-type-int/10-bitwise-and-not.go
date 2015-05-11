package main

import "fmt"

func main() {
	var x, y, andNot uint8

// START OMIT
x = 9
y = 28
andNot = x &^ y
// END OMIT

	fmt.Printf("%-6v = %.8b = %-3d\n", "x", x, x)
	fmt.Printf("%-6v = %.8b = %-3d\n", "y", y, y)
	fmt.Printf("%-6v = %.8b = %-3d", "andNot", andNot, andNot)
}
