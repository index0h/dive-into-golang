package main

import "fmt"

type Struct struct {
}

func main() {
	Bool := false
	Int := 15
	Float := float32(15.4)
	String := "A"
	Array := [1]int{1}
	Slice := make([]int, 2)
	Map := make(map[int]int, 2)
	Channel := make(chan int)
	MyStruct := Struct{}
	Pointer := &MyStruct
	Func := func () {}

// START OMIT
//_ = complex64(nil)
//_ = complex64(Bool)
//_ = complex64(Int)
//_ = complex64(Float)
//_ = complex64(String)
//_ = complex64(Array)
//_ = complex64(Slice)
//_ = complex64(Map)
//_ = complex64(Channel)
//_ = complex64(MyStruct)
//_ = complex64(Pointer)
//_ = complex64(Func)
// END OMIT

	// Used only to correct compilation
	fmt.Printf("%v (%T):\n%v\n\n", "Bool", Bool, Bool)
	fmt.Printf("%v (%T):\n%v\n\n", "Int", Int, Int)
	fmt.Printf("%v (%T):\n%v\n\n", "Float", Float, Float)
	fmt.Printf("%v (%T):\n%v\n\n", "String", String, String)
	fmt.Printf("%v (%T):\n%v\n\n", "Array", Array, Array)
	fmt.Printf("%v (%T):\n%v\n\n", "Slice", Slice, Slice)
	fmt.Printf("%v (%T):\n%v\n\n", "Map", Map, Map)
	fmt.Printf("%v (%T):\n%v\n\n", "Channel", Channel, Channel)
	fmt.Printf("%v (%T):\n%v\n\n", "Pointer", Pointer, Pointer)
	fmt.Printf("%v (%T):\n%v\n\n", "MyStruct", MyStruct, MyStruct)
	fmt.Printf("%v (%T):\n%v", "Func", Func, Func)
}
