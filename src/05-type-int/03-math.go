package main

import "fmt"

func main() {
	var sum, subtraction, product, division int

// START OMIT
sum = 2 + 3
subtraction = 10 - 2
product = 15 * 3
division = 8 / 5
// END OMIT

	fmt.Printf("%-11s (%-3T): %-5v\n", "sum", sum, sum)
	fmt.Printf("%-11s (%-3T): %-5v\n", "subtraction", subtraction, subtraction)
	fmt.Printf("%-11s (%-3T): %-5v\n", "product", product, product)
	fmt.Printf("%-11s (%-3T): %-5v", "division", division, division)
}
