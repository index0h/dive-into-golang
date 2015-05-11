package main

import "fmt"

func main() {
// START OMIT
Complex64 := complex64(1 + 2i)
Complex128 := complex128(5 - 1i)

Complex64To128 := complex128(Complex64)
Complex128To64 := complex64(Complex128)
// END OMIT

	fmt.Printf("%-14v (%-10T):\n%v\n", "Complex64", Complex64, Complex64)
	fmt.Printf("%-14v (%-10T):\n%v\n", "Complex128", Complex128, Complex128)
	fmt.Printf("%-14v (%-10T):\n%v\n", "Complex64To128", Complex64To128, Complex64To128)
	fmt.Printf("%-14v (%-10T):\n%v", "Complex128To64", Complex128To64, Complex128To64)
}
