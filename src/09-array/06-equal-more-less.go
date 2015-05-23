package main

import "fmt"

func main() {
// START OMIT
Equal := "A" == "B"
NotEqual := "A" != "B"
OneLessTwo := "A" < "B"
OneMoreTwo := "A" > "B"
OneSoftLessTwo := "B" <= "B"
OneSoftMoreTwo := "B" >= "B"
// END OMIT

	fmt.Printf("%-14s (%-4T): %v\n", "Equal", Equal, Equal)
	fmt.Printf("%-14s (%-4T): %v\n", "NotEqual", NotEqual, NotEqual)
	fmt.Printf("%-14s (%-4T): %v\n", "OneLessTwo", OneLessTwo, OneLessTwo)
	fmt.Printf("%-14s (%-4T): %v\n", "OneMoreTwo", OneMoreTwo, OneMoreTwo)
	fmt.Printf("%-14s (%-4T): %v\n", "OneSoftLessTwo", OneSoftLessTwo, OneSoftLessTwo)
	fmt.Printf("%-14s (%-4T): %v", "OneSoftMoreTwo", OneSoftMoreTwo, OneSoftMoreTwo)
}
