package main

import "fmt"

func main() {
// START OMIT
Int8 := int8(127)
Uint8To8 := uint8(Int8)
Uint8To16 := uint16(Int8)

Int16 := int16(16383)
Uint16To8 := uint8(Int16)
Uint16To16 := uint16(Int16)
// END OMIT

	fmt.Printf("%-10v: %.16b\n", "Int8", Int8)
	fmt.Printf("%-10v: %.16b\n", "Uint8To8", Uint8To8)
	fmt.Printf("%-10v: %.16b\n", "Uint8To16", Uint8To16)

	fmt.Printf("%-10v: %.16b\n", "Int16", Int16)
	fmt.Printf("%-10v: %.16b\n", "Uint16To8", Uint16To8)
	fmt.Printf("%-10v: %.16b", "Uint16To16", Uint16To16)
}
