package main

import "fmt"

type Struct struct {
}

func main() {
	Bool := false
	Complex := complex64(3 + 2i)
	String := "A"
	Array := [1]int{1}
	Slice := make([]int, 2)
	Map := make(map[int]int, 2)
	Channel := make(chan int)
	MyStruct := Struct{}
	Pointer := &MyStruct
	Func := func () {}

// START OMIT
//_ = float32(nil)
//_ = float32(Bool)
//_ = float32(Complex)
//_ = float32(String)
//_ = float32(Array)
//_ = float32(Slice)
//_ = float32(Map)
//_ = float32(Channel)
//_ = float32(MyStruct)
//_ = float32(Pointer)
//_ = float32(Func)
// END OMIT

	// Used only to correct compilation
	fmt.Printf("%v (%T):\n%v\n\n", "Bool", Bool, Bool)
	fmt.Printf("%v (%T):\n%v\n\n", "Complex", Complex, Complex)
	fmt.Printf("%v (%T):\n%v\n\n", "String", String, String)
	fmt.Printf("%v (%T):\n%v\n\n", "Array", Array, Array)
	fmt.Printf("%v (%T):\n%v\n\n", "Slice", Slice, Slice)
	fmt.Printf("%v (%T):\n%v\n\n", "Map", Map, Map)
	fmt.Printf("%v (%T):\n%v\n\n", "Channel", Channel, Channel)
	fmt.Printf("%v (%T):\n%v\n\n", "Pointer", Pointer, Pointer)
	fmt.Printf("%v (%T):\n%v\n\n", "MyStruct", MyStruct, MyStruct)
	fmt.Printf("%v (%T):\n%v", "Func", Func, Func)
}
