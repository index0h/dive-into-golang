package main

import "fmt"

func main() {
// START OMIT
var First = [...][2]byte{{1}, {3, 4}}
Second := [2][2][2]byte{}
// END OMIT

	fmt.Printf("%-6s (%-7T):\n%v\n\n", "First", First, First)
	fmt.Printf("%-6s (%-7T):\n%v", "Second", Second, Second)
}
