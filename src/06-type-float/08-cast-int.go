package main

import "fmt"

func main() {
// START OMIT
Int8 := int8(127)
Float32 := float32(Int8)
Float64 := float64(Int8)
// END OMIT

	fmt.Printf("%-7v (%-7T): %v\n", "Int8", Int8, Int8)
	fmt.Printf("%-7v (%-7T): %.2f\n", "Float32", Float32, Float32)
	fmt.Printf("%-7v (%-7T): %.2f", "Float64", Float64, Float64)
}
