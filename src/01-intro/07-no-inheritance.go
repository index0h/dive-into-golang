package main

import "fmt"

// START OMIT
type Embedded struct {
	First string
}
type Aggregated Embedded
type Parent struct {
	Embedded
	Aggr Aggregated
}

func main() {
	Var := Parent{}
	Var.First = "first"
	Var.Aggr.First = "second"
	fmt.Println(Var)
}
// END OMIT
