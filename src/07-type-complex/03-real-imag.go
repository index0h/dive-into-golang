package main

import "fmt"

func main() {
// START OMIT
Complex := complex(2.1, 2)
Real := real(Complex)
Img := imag(Complex)
// END OMIT

	fmt.Printf("%-7s (%-10T): %v\n", "Complex", Complex, Complex)
	fmt.Printf("%-7s (%-10T): %v\n", "Real", Real, Real)
	fmt.Printf("%-7s (%-10T): %v", "Img", Img, Img)
}
