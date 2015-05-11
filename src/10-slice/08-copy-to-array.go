package main

import "fmt"

func main() {
// START OMIT
Array := [3]int{}
Slice := []int{1, 2, 3, 4, 5, 6}

Count := copy(Array[:], Slice)

Slice[0] *= 5
Slice[1] *= 5
Slice[2] *= 5
Slice[3] *= 5
Slice[4] *= 5
Slice[5] *= 5
// END OMIT

	fmt.Printf("%-6s (%T): value: %v\n\n", "Count", Count, Count)
	fmt.Printf("%-6s (%T):  len: %v cap: %v\nvalue: %v\n\n", "Slice", Slice, len(Slice), cap(Slice), Slice)
	fmt.Printf("%-6s (%T): len: %v cap: %v\nvalue: %v", "Array", Array, len(Array), cap(Array), Array)
}
