package main

import (
	"fmt"
)

func main() {
	first := make(map[bool]float32, 20)
	second := make(map[string]bool)
	third := map[string]uint{"one": 1, "two": 2, "three": 3}
	fourth := map[string]uint{}

	fmt.Println("first  =", first)  // first  = map[]
	fmt.Println("second =", second) // second = map[]
	fmt.Println("third  =", third)  // third  = map[one:1 two:2 three:3]
	fmt.Println("fourth =", fourth) // fourth = map[]

	fmt.Println("\n------------------------------------------------------------------\n")

	first[false] = 100
	first[true] = 200

	second["true"] = true
	second["false"] = false

	third["four"] = 4
	third["five"] = 5

	fmt.Println("first  =", first)  // first  = map[false:100 true:200]
	fmt.Println("second =", second) // second = map[true:true false:false]
	fmt.Println("third  =", third)  // third  = map[one:1 two:2 three:3 four:4 five:5]

	fmt.Println("\n------------------------------------------------------------------\n")

	delete(third, "three")
	delete(third, "two")

	fmt.Println("third =", third) // third = map[five:5 one:1 four:4]
}
