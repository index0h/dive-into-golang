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

	fmt.Printf("%-5s (%-6T): len: %v cap: %v\n", "Array", Array, len(Array), cap(Array))
	fmt.Printf("%-5s (%-6T): len: %v cap: %v\n", "Part1", Part1, len(Part1), cap(Part1))
	fmt.Printf("%-5s (%-6T): len: %v cap: %v\n", "Part2", Part2, len(Part2), cap(Part2))
	fmt.Printf("%-5s (%-6T): len: %v cap: %v\n", "Part3", Part3, len(Part3), cap(Part3))
	fmt.Printf("%-5s (%-6T): len: %v cap: %v", "Full", Full, len(Full), cap(Full))
}
