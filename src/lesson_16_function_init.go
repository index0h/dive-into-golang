package main

import (
	"fmt"
	_ "lesson_16_function_init"
)

func init () {
	fmt.Println("Call: main.init()")
}

func main () {
	fmt.Println("Call: main.main()")
}
