package main

import "fmt"

/// START OMIT
type EmptyChecker interface {
	Check(string) bool
}
type Implementer struct {
}
func (imp Implementer) Check(data string) bool {
	return data == ""
}
func ImWorkingWithChecker(checker EmptyChecker) {
	fmt.Println(checker.Check(""))
}

func main() {
	ImWorkingWithChecker(Implementer{})
}
// END OMIT
