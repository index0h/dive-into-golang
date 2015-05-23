package main

import "fmt"

func main() {
// START OMIT
MyFunc := func(value string) {
	fmt.Println(value)
}

MyFunc("Hello world")
// END OMIT
}
