package main

import (
	"fmt"
	"os"
)

func main() {
	trueVar := true
	falseVar := false

	fmt.Print("First:  ")
	if trueVar {
		fmt.Println("if trueVar") // First:  if trueVar
	}

	fmt.Print("Second: ")
	if falseVar {
		fmt.Println("if falseVar")
	} else {
		fmt.Println("if falseVar -> else") // Second: if falseVar -> else
	}

	fmt.Print("Third:  ")
	if falseVar {
		fmt.Println("if falseVar")
	} else if trueVar {
		fmt.Println("if falseVar -> if trueVar") // Third:  if falseVar -> if trueVar
	} else {
		fmt.Println("if falseVar -> if trueVar -> else")
	}

	fmt.Print("Fourth: ")
	if _, err := os.Stat("some unknown file"); err == nil {
		fmt.Println("unknown file exists")
	} else if _, err := os.Stat("/home/vagrant/work/src/lesson_12_instruction_if.go"); err == nil {
		fmt.Println("real file exists") // Fourth: real file exists
	} else {
		fmt.Println("both files not exist")
	}

}
