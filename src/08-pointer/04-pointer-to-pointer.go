package main

import "fmt"

func main() {
// START OMIT
Pointer1 := new(byte)
Pointer2 := &Pointer1

**Pointer2 = 50
// END OMIT
	fmt.Printf("%s (%T):\n*value: %v\nvalue : %v\naddr  : %p\n\n", "Pointer1", Pointer1, *Pointer1, Pointer1, &Pointer1)
	fmt.Printf(
		"%s (%T):\nvalue  : %v\n*value : %v\n**value: %v\naddr   : %p",
		"Pointer2",
		Pointer2,
		Pointer2,
		*Pointer2,
		**Pointer2,
		&Pointer2,
	)
}
