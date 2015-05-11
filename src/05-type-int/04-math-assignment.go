package main

import "fmt"

func main() {
	var sum, subtraction, product, division int

// START OMIT
sum = 2
sum += 8

subtraction = 15
subtraction -= 7

product = 9
product *= 2

division = 8
division /= 4
// END OMIT

	fmt.Printf("%-11s (%-3T): %-5v\n", "sum", sum, sum)
	fmt.Printf("%-11s (%-3T): %-5v\n", "subtraction", subtraction, subtraction)
	fmt.Printf("%-11s (%-3T): %-5v\n", "product", product, product)
	fmt.Printf("%-11s (%-3T): %-5v", "division", division, division)
}
