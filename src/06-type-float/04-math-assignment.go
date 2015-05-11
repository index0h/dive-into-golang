package main

import "fmt"

func main() {
	var sum, subtraction, product, division float32

// START OMIT
sum = 2.1
sum += 8.4

subtraction = 15.5
subtraction -= 7.5

product = 9.3
product *= 2.2

division = 8.4
division /= 4.1
// END OMIT

	fmt.Printf("%-11s (%T): %-5.2f\n", "sum", sum, sum)
	fmt.Printf("%-11s (%T): %-5.2f\n", "subtraction", subtraction, subtraction)
	fmt.Printf("%-11s (%T): %-5.2f\n", "product", product, product)
	fmt.Printf("%-11s (%T): %-5.2f", "division", division, division)
}
