package main

import (
	"fmt"
)

func main() {
	var first [2]bool
	second := [...]int {-1, -2, -3}
	third := [6]uint {1, 2, 3, 4, 5, 6}

	var fourth [2][2]float32
	fifth := [2][2]float32 {{1}, {5.4, 7.7}}

	fmt.Printf("first  len(%d) = %v\n",   len(first),  first)  // first  len(2) = [false false]
	fmt.Printf("second len(%d) = %v\n",   len(second), second) // second len(3) = [-1 -2 -3]
	fmt.Printf("third  len(%d) = %v\n",   len(third),  third)  // third  len(6) = [1 2 3 4 5 6]
	fmt.Printf("fourth len(%d) = %v\n",   len(fourth), fourth) // fourth len(2) = [[0 0] [0 0]]
	fmt.Printf("fifth  len(%d) = %v\n\n", len(fifth),  fifth)  // fifth  len(2) = [[1 0] [5.4 7.7]]

	fmt.Println("second[2]   = ", second[2])   // second[2]   = -3
	fmt.Println("fifth[0][1] = ", fifth[1][0]) // fifth[0][1] = 5.4

	fmt.Println("\n------------------------------------------------------------------\n")

	copyOfThird := third
	third[4] = 500
	third[5] = 600
	fmt.Println("third       = ", third)       // third       = [1 2 3 4 500 600]
	fmt.Println("copyOfThird = ", copyOfThird) // copyOfThird = [1 2 3 4 5 6]
}
