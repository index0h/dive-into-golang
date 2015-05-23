package main

import "fmt"

func main() {
// START OMIT
var VarUint uint
var VarBool bool
// END OMIT

	fmt.Printf("%s (%T): %v\n", "VarUint", VarUint, VarUint)
	fmt.Printf("%s (%T): %v", "VarBool", VarBool, VarBool)
}
