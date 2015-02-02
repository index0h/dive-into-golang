package main

import (
	"fmt"
)

func main() {
	first := make([]byte, 5, 15)
	second := make([]byte, 3)
	var third []bool
	fourth := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("first  len(%d) cap(%d) = %v\n", len(first), cap(first),  first)   // first  len(5) cap(15) = [0 0 0 0 0]
	fmt.Printf("second len(%d) cap(%d) = %v\n", len(second), cap(second), second) // second len(7) cap(3)  = [0 0 0]
	fmt.Printf("third  len(%d) cap(%d) = %v\n", len(third), cap(third),  third)   // third  len(0) cap(0)  = []
	fmt.Printf("fourth len(%d) cap(%d) = %v\n", len(fourth), cap(fourth), fourth) // fourth len(6) cap(6) = [1 2 3 4 5 6]

	fmt.Println("\n------------------------------------------------------------------\n")

	part1 := fourth[:2]
	part2 := fourth[2:4]
	part3 := fourth[4:]
	part4 := fourth[:]

	fmt.Printf("part1 len(%d) cap(%d) = %v\n", len(part1), cap(part1), part1) // part1 len(2) cap(6) = [1 2]
	fmt.Printf("part2 len(%d) cap(%d) = %v\n", len(part2), cap(part2), part2) // part2 len(2) cap(4) = [3 4]
	fmt.Printf("part3 len(%d) cap(%d) = %v\n", len(part3), cap(part3), part3) // part3 len(2) cap(2) = [5 6]
	fmt.Printf("part4 len(%d) cap(%d) = %v\n", len(part4), cap(part4), part4) // part4 len(6) cap(6) = [1 2 3 4 5 6]

	fmt.Println("\n------------------------------------------------------------------\n")

	part1[0] = 100
	part2[0] = 300
	part3[0] = 500

	fmt.Println("fourth = ", fourth) // fourth = [100 2 300 4 500 6]
	fmt.Println("part4  = ", part4)  // part4  = [100 2 300 4 500 6]

	fmt.Println("\n------------------------------------------------------------------\n")

	copyOfFourth := make([]int, 6)
	count := copy(copyOfFourth, fourth)
	fourth = []int {1, 2, 3, 4, 5, 6}

	fmt.Println("fourth        = ", fourth)        // fourth        = [1 2 3 4 5 6]
	fmt.Println("copyOfFourth  = ", copyOfFourth)  // copyOfFourth  = [100 2 300 4 500 6]
	fmt.Println("count         = ", count)         // count         = 6

	fmt.Println("\n------------------------------------------------------------------\n")

	fourth = append(fourth, copyOfFourth[:3]...)

	fmt.Println("fourth        = ", fourth)        // fourth        = [1 2 3 4 5 6 100 2 300]
	fmt.Println("copyOfFourth  = ", copyOfFourth)  // copyOfFourth  = [100 2 300 4 500 6]

	fourth = append(fourth, 111, 222, 333)
	fmt.Println("fourth        = ", fourth)        // fourth        = [1 2 3 4 5 6 100 2 300 111 222 333]

	fmt.Println("\n------------------------------------------------------------------\n")

	fmt.Printf("fourth len(%d) cap(%d)\n", len(fourth), cap(fourth))
	fourth = append(fourth, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("fourth len(%d) cap(%d)\n", len(fourth), cap(fourth))
}

