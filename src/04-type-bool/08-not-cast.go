package main

import "fmt"

type Struct struct {
}

func main() {
	Int := 15
	Float := float32(15.4)
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
//_ = bool(nil)
//_ = bool(Int)
//_ = bool(Complex)
//_ = bool(String)
//_ = bool(Array)
//_ = bool(Slice)
//_ = bool(Map)
//_ = bool(Channel)
//_ = bool(MyStruct)
//_ = bool(Pointer)
//_ = bool(Func)
// END OMIT

	// Used only to correct compilation
	fmt.Printf("%v (%T):\n%v\n\n", "Int", Int, Int)
	fmt.Printf("%v (%T):\n%v\n\n", "Float", Float, Float)
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
