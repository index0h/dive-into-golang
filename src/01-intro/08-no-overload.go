package main

import "fmt"

/// START OMIT
type Embedded struct {
	First string
}
type Parent struct {
	Embedded
	First string
}

func main() {
	Var := Parent{}
	Var.First = "first"
	Var.Embedded.First = "second"
	fmt.Println(Var)
}
// END OMIT
