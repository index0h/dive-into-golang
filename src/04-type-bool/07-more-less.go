package main

import "fmt"

func main() {
// START OMIT
First := 1 < 2
Second := 1 > 2
Third := 2 <= 2
Fourth := 2 >= 2
// END OMIT

	fmt.Printf("%-6s (%-4T): %-5v\n", "First", First, First)
	fmt.Printf("%-6s (%-4T): %-5v\n", "Second", Second, Second)
	fmt.Printf("%-6s (%-4T): %-5v\n", "Third", Third, Third)
	fmt.Printf("%-6s (%-4T): %-5v", "Fourth", Fourth, Fourth)
}
