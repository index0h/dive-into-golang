package main

import (
	"fmt"
)

func main() {
	first := make(map[bool]float32, 20)
	second := make(map[string]bool)
	third := map[string]uint{"one": 1, "two": 2, "three": 3}
	fourth:= map[string]uint{}

	fmt.Println("first  =", first)
	fmt.Println("second =", second)
	fmt.Println("third  =", third)
	fmt.Println("fourth =", fourth)

	fmt.Println("\n------------------------------------------------------------------\n")

	first[false] = 100
	first[true] = 200

	second["true"] = true
	second["false"] = false

	third["four"] = 4
	third["five"] = 5

	fmt.Println("first  =", first)
	fmt.Println("second =", second)
	fmt.Println("third  =", third)

	fmt.Println("\n------------------------------------------------------------------\n")

	delete(third, "three")
	delete(third, "two")

	fmt.Println("third =", third)
}
