package main

import (
	"fmt"
)

func main() {

	fmt.Print("First:  ")
	switch {
	case 1 < 2:
		fmt.Println("1 < 2") // First:  TRUE  1 < 2
	case 2 > 3:
		fmt.Println("2 > 3")
	}

	fmt.Print("Second: ")
	switch suffix := "jpg"; suffix {
	case "png":
		fallthrough
	case "jpg":
		fallthrough
	case "gif":
		fmt.Println("Image") // Second: Image
	default:
		fmt.Println("Not image")
	}

	var someObject interface{}
	someObject = 5

	fmt.Print("Third:  ")
	switch someObject.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Sorry, I don't know") // Third:  Sorry, I don't know
	}

}
