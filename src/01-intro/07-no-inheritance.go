package main

import "fmt"

// START STRUCT OMIT
type SubStruct struct {
	Data string
}

type Parent struct {
	Data string
	SubStruct
	Aggregated SubStruct
}
// END STRUCT OMIT

func main() {
// START OMIT
Var := Parent{}
Var.Data = "first"
Var.SubStruct.Data = "second"
Var.Aggregated.Data = "third"
fmt.Println(Var)
// END OMIT
}
