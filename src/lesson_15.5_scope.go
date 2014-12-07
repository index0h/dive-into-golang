package main

import "fmt"

var first, second, format string

func main () {
	// Working with package-level variables
	first, second, format = "first", "second", "first = %-16v | second = %v\n"
	fmt.Printf(format, first, second)         // first = first            | second = second

	// Redefine variable by if instruction
	if first := "if"; first != "" {
		fmt.Printf(format, first, second)     // first = if               | second = second
	}

	// Redefine variable by switch instruction
	switch first := "switch"; first {
	case "switch":
		fmt.Printf(format, first, second)     // first = switch           | second = second
	default:
		first = "switch-default"
		fmt.Printf(format, first, second)
	}

	// Redefine variable by for instruction
	for first := 0; first < 1; first++ {
		fmt.Printf(format, first, second)     // first = 0                | second = second
	}

	// Working with package-level variables
	func () {
		fmt.Printf(format, first, second)     // first = first            | second = second
	}()

	// Redefine variable inside function
	func () {
		first := "func"
		fmt.Printf(format, first, second)     // first = func             | second = second

		first = "func-in-func"
		func () {
			fmt.Printf(format, first, second) // first = func-in-func     | second = second
		}()
	}()

	// Redefine variable inside braces
	{
		first := "braces"
		fmt.Printf(format, first, second)     // first = braces           | second = second

		first = "braces-in-braces"
		{
			fmt.Printf(format, first, second) // first = braces-in-braces | second = second
		}
	}

	// Working with package-level variables
	fmt.Printf(format, first, second)         // first = first            | second = second
}
