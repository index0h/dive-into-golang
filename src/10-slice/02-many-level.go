package main

import "fmt"

func main() {
// START OMIT
var First = [][]byte{{1}, {3, 4}}
Second := [][][]byte{}
// END OMIT

	fmt.Printf("%-6s (%-7T):\n%v\n\n", "First", First, First)
	fmt.Printf("%-6s (%-7T):\n%v", "Second", Second, Second)
}
