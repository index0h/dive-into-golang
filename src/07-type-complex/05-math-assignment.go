package main

import "fmt"

func main() {
	var sum, subtraction, product, division complex64

// START OMIT
sum = 2 + 1i
sum += 8 -3i

subtraction = 15 + 2i
subtraction -= 7 + 1i

product = 9
product *= 2 + 2i

division = 8
division /= 4 -2i
// END OMIT

	fmt.Printf("%-11s: %v\n", "sum", sum, )
	fmt.Printf("%-11s: %v\n", "subtraction", subtraction)
	fmt.Printf("%-11s: %v\n", "product", product)
	fmt.Printf("%-11s: %v", "division", division)
}
