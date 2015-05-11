package main

import (
	"fmt"
	my "github.com/index0h/dive-into-golang/src/01-intro/public_private"
)

func main() {
// START OMIT
data := my.MyStruct{}
data.Public = "I'm public"
data.private = "I'm private"
// END OMIT

	fmt.Printf("data (%T):\n%v", data, data)
}
