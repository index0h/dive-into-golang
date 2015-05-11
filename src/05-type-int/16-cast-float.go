package main

import "fmt"

func main() {
// START OMIT
Float := float32(127.148)
Int := int(Float)
// END OMIT

	fmt.Printf("%-5v (%-7T) = %v\n", "Float", Float, Float)
	fmt.Printf("%-5v (%-7T) = %v", "Int", Int, Int)
}
