package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var defaultFloat float32
	floatVar := 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)                     // 0
	fmt.Printf("floatVar (%s) = %v\n", reflect.TypeOf(floatVar), floatVar) // 5.5

	fmt.Println("\n------------------------------------------------------------------\n")

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")
}
