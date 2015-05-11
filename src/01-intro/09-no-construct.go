package main

import "fmt"

/// START OMIT
type MyStruct struct {
	First string
}

func NewMyStruct(first string) *MyStruct {
	return &MyStruct{First: first}
}

func main() {
	Var := NewMyStruct("Hello world")
	fmt.Println(Var)
}
// END OMIT
