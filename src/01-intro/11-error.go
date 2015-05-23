package main

import (
	"fmt"
	"errors"
)

/// START OMIT
func ErrorFunc() error {
	return errors.New("Something went wrong")
}
func RecoverFunc() {
	fmt.Println("There was:", recover())
}

func main() {
	if err := ErrorFunc(); err != nil {
		fmt.Println(err)
	}

	defer RecoverFunc()
	panic("BOOM!")
}
// END OMIT
