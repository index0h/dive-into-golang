package main

import (
	"fmt"
)

func main() {
	first := 10
	second := 10
	third := []string{"a", "b", "c", "d", "e", "f"}
	fourth := map[string]string{"a": "A", "b": "B", "c":"C", "d":"D", "e":"E", "f":"F"}


	fmt.Println("First: first =", first)
	for {
		fmt.Print(".")
		if first <= 0 {
			break
		}
		first--
	}
	fmt.Println("\nfirst =", first)

	fmt.Println("\n------------------------------------------------------------------\n")

	fmt.Println("Second: second =", second)
	for second = 10; second < 20; second++ {
		fmt.Print(".")
	}
	fmt.Println("\nsecond =", second)

	fmt.Println("\n------------------------------------------------------------------\n")

	fmt.Println("Third: third =", third)
	for index := range third {
		fmt.Println(index, " -> ", third[index])
	}

	fmt.Println("\n------------------------------------------------------------------\n")

	fmt.Println("Fourth: fourth =", fourth)
	for key, value := range third {
		fmt.Println(key, " -> ", value)
	}
}
