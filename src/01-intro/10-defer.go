package main

import (
	"fmt"
)

/// START OMIT
func main() {
	defer func() {
		fmt.Println("main: defer")
	}()

	fmt.Println("main: end")
}
// END OMIT
