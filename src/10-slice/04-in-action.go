package main

import "fmt"

func main() {
// START OMIT
Array := [...]int{1, 2, 3, 4, 5, 6}

Part1 := Array[:2]
Part2 := Array[2:4]
Part3 := Array[4:]
Full := Array[:]
// END OMIT

	fmt.Printf("%-5s (%-6T): %v\n", "Array", Array, Array)
	fmt.Printf("%-5s (%-6T): %v\n", "Part1", Part1, Part1)
	fmt.Printf("%-5s (%-6T): %v\n", "Part2", Part2, Part2)
	fmt.Printf("%-5s (%-6T): %v\n", "Part3", Part3, Part3)
	fmt.Printf("%-5s (%-6T): %v", "Full", Full, Full)
}
