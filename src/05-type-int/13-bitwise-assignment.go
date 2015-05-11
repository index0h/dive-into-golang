package main

import "fmt"

func main() {
	var y, and, or, xor, andNot, left, right, mod int

// START OMIT
y = 28
and, or, xor, andNot, left, right, mod = 9, 9, 9, 9, 9, 9, 9

and &= y
or |= y
xor ^= y
andNot &^= y
left <<= 2
right >>= 2
mod %= y
// END OMIT

	fmt.Printf("%-6v = %.8b = %-3d\n", "y", y, y)
	fmt.Printf("%-6v = %.8b = %-3d\n", "and", and, and)
	fmt.Printf("%-6v = %.8b = %-3d\n", "or", or, or)
	fmt.Printf("%-6v = %.8b = %-3d\n", "xor", xor, xor)
	fmt.Printf("%-6v = %.8b = %-3d\n", "andNot", andNot, andNot)
	fmt.Printf("%-6v = %.8b = %-3d\n", "left", left, left)
	fmt.Printf("%-6v = %.8b = %-3d\n", "right", right, right)
	fmt.Printf("%-6v = %.8b = %-3d", "mod", mod, mod)
}
