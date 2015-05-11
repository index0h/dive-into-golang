package main

import "fmt"

func main() {
// START OMIT
var EmptyPointer *byte

IsNil := EmptyPointer == nil
// END OMIT

	fmt.Printf("%s (%T)\nvalue: %v\naddr : %p\n", "EmptyPointer", EmptyPointer, EmptyPointer, EmptyPointer)
	fmt.Printf("isNil: %v", IsNil)
}
