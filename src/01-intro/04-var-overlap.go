package main

import "fmt"

func main() {
// START OMIT
i := 10

if i := 2; i > 1 {
	fmt.Printf("%-5v = %v\n", "IF: i", i)
}

fmt.Printf("%-5v = %v", "i", i)
// END OMIT
}
