package main

import "fmt"

func main() {
// START OMIT
Float32 := float32(142.8)
Float64 := float64(544.2)

Float32To64 := float64(Float32)
Float64To32 := float32(Float64)
// END OMIT

	fmt.Printf("%-11v (%T): %.1f\n", "Float32", Float32, Float32)
	fmt.Printf("%-11v (%T): %.1f\n", "Float64", Float64, Float64)
	fmt.Printf("%-11v (%T): %.1f\n", "Float32To64", Float32To64, Float32To64)
	fmt.Printf("%-11v (%T): %.1f", "Float64To32", Float64To32, Float64To32)
}
