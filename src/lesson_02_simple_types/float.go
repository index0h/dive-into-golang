package lesson_02_simple_types

import (
	"fmt"
	"math"
)

func DemonstrateFloat() {
	fmt.Println("\t------------------------------------------------------------------\n\n")


	fmt.Println("\tПределы вещественных чисел\n")

	fmt.Println("MAX float32 = ", math.MaxFloat32)
	fmt.Println("MIN float32 = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64 = ", math.MaxFloat64)
	fmt.Println("MIN float64 = ", math.SmallestNonzeroFloat64, "\n")
}
