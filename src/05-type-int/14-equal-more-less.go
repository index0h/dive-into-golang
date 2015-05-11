package main

import "fmt"

func main() {
// START OMIT
Equal := 1 == 2
NotEqual := 1 != 2
OneLessTwo := 1 < 2
OneMoreTwo := 1 > 2
OneSoftLessTwo := 2 <= 2
OneSoftMoreTwo := 2 >= 2
// END OMIT

	fmt.Printf("%-14s (%-4T): %v\n", "Equal", Equal, Equal)
	fmt.Printf("%-14s (%-4T): %v\n", "NotEqual", NotEqual, NotEqual)
	fmt.Printf("%-14s (%-4T): %v\n", "OneLessTwo", OneLessTwo, OneLessTwo)
	fmt.Printf("%-14s (%-4T): %v\n", "OneMoreTwo", OneMoreTwo, OneMoreTwo)
	fmt.Printf("%-14s (%-4T): %v\n", "OneSoftLessTwo", OneSoftLessTwo, OneSoftLessTwo)
	fmt.Printf("%-14s (%-4T): %v", "OneSoftMoreTwo", OneSoftMoreTwo, OneSoftMoreTwo)
}
