package main

import "fmt"

func main() {
// START OMIT
var First [2]byte
var Second [2][2]byte

First[0] = 1
First[1] = 2

Second[0][0] = 1
Second[0][1] = 2
Second[1][0] = 3
Second[1][1] = 4
// END OMIT

	fmt.Printf("%-6s (%T):\n%v\n\n", "First", First, First)
	fmt.Printf("%-6s (%T):\n%v", "Second", Second, Second)
}
