package main

import "fmt"

func main() {
// START OMIT
var Byte byte = 10

Pointer := &Byte
*Pointer = 25
// END OMIT

	fmt.Printf("%s (%T):\nvalue: %v\naddr : %p\n\n", "Byte", Byte, Byte, &Byte)
	fmt.Printf("%s (%T):\nvalue : %v\n*value: %v\naddr  : %p", "Pointer", Pointer, Pointer, *Pointer, &Pointer)
}
