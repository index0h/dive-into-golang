package main

import "fmt"

// START OMIT
type DuckInterface interface {
	Quack() string
}
type ImDuck struct {
}
func (duck ImDuck) Quack() string {
	return "quack-quack"
}
func ImWorkingWithDuck(duck DuckInterface) {
	fmt.Printf("%T: %v", duck, duck.Quack())
}

func main() {
	ImWorkingWithDuck(ImDuck{})
}
// END OMIT
