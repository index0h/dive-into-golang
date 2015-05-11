package main

import "fmt"

func main() {
// START OMIT
Equal := 1.5 + 1i == 2.5 + 2i
// Equal = 2.5 + 2i == (2.5 + 2i + 0.01 * 2.1 - 0.021)
NotEqual := 1.5 + 1i != 2.5 + 2i
// END OMIT

	fmt.Printf("%-8s (%-4T): %v\n", "Equal", Equal, Equal)
	fmt.Printf("%-8s (%-4T): %v", "NotEqual", NotEqual, NotEqual)
}
