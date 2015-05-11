package main

import "fmt"

func main() {
	var sum, subtraction, product, division complex64

// START OMIT
sum = 2.1 +2i + 3 + 5i
subtraction = 1.05 - 2 -1i
product = 15 * 3.12
division = 8 / 5 // !!! int operation
//division = 8.0 / 5
// END OMIT

	fmt.Printf("%-11s: %v\n", "sum", sum)
	fmt.Printf("%-11s: %v\n", "subtraction", subtraction)
	fmt.Printf("%-11s: %v\n", "product", product)
	fmt.Printf("%-11s: %v", "division", division)
}
