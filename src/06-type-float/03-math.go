package main

import "fmt"

func main() {
	var sum, subtraction, product, division float32

// START OMIT
sum = 2.1 + 3
subtraction = 1.05 - 2
product = 15 * 3.12
division = 8 / 5 // !!! int operation
//division = 8.0 / 5
// END OMIT

	fmt.Printf("%-11s (%-3T): %-5v\n", "sum", sum, sum)
	fmt.Printf("%-11s (%-3T): %-5v\n", "subtraction", subtraction, subtraction)
	fmt.Printf("%-11s (%-3T): %-5v\n", "product", product, product)
	fmt.Printf("%-11s (%-3T): %-5v", "division", division, division)
}
