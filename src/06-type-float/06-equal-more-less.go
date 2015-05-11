package main

import "fmt"

func main() {
// START OMIT
Equal := 1.5 == 2.5
// Equal = 2.5 == (2.5 + 0.01 * 2.1 - 0.021)
NotEqual := 1.5 != 2.5
OneLessTwo := 1.5 < 2.5
OneMoreTwo := 1.5 > 2.5
OneSoftLessTwo := 2.5 <= 2.5
OneSoftMoreTwo := 2.5 >= 2.5
// END OMIT

	fmt.Printf("%-14s (%-4T): %v\n", "Equal", Equal, Equal)
	fmt.Printf("%-14s (%-4T): %v\n", "NotEqual", NotEqual, NotEqual)
	fmt.Printf("%-14s (%-4T): %v\n", "OneLessTwo", OneLessTwo, OneLessTwo)
	fmt.Printf("%-14s (%-4T): %v\n", "OneMoreTwo", OneMoreTwo, OneMoreTwo)
	fmt.Printf("%-14s (%-4T): %v\n", "OneSoftLessTwo", OneSoftLessTwo, OneSoftLessTwo)
	fmt.Printf("%-14s (%-4T): %v", "OneSoftMoreTwo", OneSoftMoreTwo, OneSoftMoreTwo)
}
