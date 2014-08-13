package lesson_02_simple_types

import (
	"fmt"
	"reflect"
)

func DemonstrateComplex() {
	first := 1.5 + 2.78i
	second := complex64(1.3 + 17i)
	third := complex(15, 1.7) // complex128(15 + 1.7i)

	fmt.Println("\t------------------------------------------------------------------\n\n")


	fmt.Println("\tКомплексные числа\n")

	fmt.Printf("first  (%s) = %v\n", reflect.TypeOf(first), first)
	fmt.Printf("second (%s)  = %v\n", reflect.TypeOf(second), second)
	fmt.Printf("third  (%s) = %v\n", reflect.TypeOf(third), third)
}
