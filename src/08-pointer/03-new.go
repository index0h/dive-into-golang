package main

import "fmt"

func main() {
// START OMIT
Pointer := new(byte)

*Pointer = 15
// END OMIT
	fmt.Printf("%s (%T):\nvalue : %v\n*value: %v\naddr  : %p", "Pointer", Pointer, Pointer, *Pointer, &Pointer)
}
