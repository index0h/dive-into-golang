package main

import "fmt"

func main() {
// START OMIT
Float32 := float32(1)

Complex64 := complex(Float32, Float32)
// Complex64 = complex(int(1), int(1)) // Error
// END OMIT
	fmt.Printf("%-9v (%-9T) = %v\n", "Float32", Float32, Float32)
	fmt.Printf("%-9v (%-9T) = %v", "Complex64", Complex64, Complex64)
}
