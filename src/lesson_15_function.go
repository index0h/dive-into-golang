package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Example of simple function with two arguments
func print(prefix, data string) {
	fmt.Printf("Function %-20v: %v\n", prefix , data)
}

// Example of variable count of arguments
func concatenation(arguments ...string) string {
	result := ""
	for _, part := range arguments {
		result += part
	}

	return result
}

// Example of recursive function
//
// pow(2, 5) > pow(2, 4) > pow(2, 3) > pow(2, 2) > pow(2, 1)
// 2 * 16    < 2 * 8     < 2 * 4     < 2 * 2     < 2
func pow(number int, degree uint) int {
	if degree <= 1 {
		return number
	}

	return number * pow(number, degree - 1)
}

// Example of function as result of another function
func wrap(wrapSymbols string) (func(data string) string) {
	return func(data string) string {
		return wrapSymbols + data + wrapSymbols
	}
}

// Example of using functions as arguments
func stringConverter(data string, middlewareList ...func (data string) string) string {
	for _, middleware := range middlewareList {
		data = middleware(data)
	}

	return data
}

func main () {
	print("print", "some data")

	print("concatenation", concatenation("a", "b", "c", "d"))

	aliasOfConcatenation := concatenation

	print("aliasOfConcatenation", aliasOfConcatenation("alias", " ", "of", " ", "concatenation"))

	print("pow", strconv.Itoa(pow(2, 5)))

	print("wrap", wrap("|")("wrap me"))

	print("stringConverter", stringConverter("	\n\t SoMe tExT\n\r\t  ", strings.ToLower, strings.TrimSpace, wrap("@")))
}
