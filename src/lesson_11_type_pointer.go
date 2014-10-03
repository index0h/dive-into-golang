package main

import (
	"fmt"
)

// Переменные, хранящие ссылки:
//
// * срезы
// * отображения
// * каналы
// * функции
// * указатели

func main() {
	intVar := 10

	firstPointer := &intVar
	secondPointer := &firstPointer
	var thirdPointer *int

	tableRowFormat := "| %-15v | %-8T | %-12p | %-12v |\n"

	fmt.Println("| Название        | Тип      | Адрес        | Значение     |")
	fmt.Println("|-----------------|----------|--------------|--------------|")

	fmt.Printf(tableRowFormat, "intvar", intVar, &intVar, intVar)
	fmt.Println("|-----------------|----------|--------------|--------------|")
	fmt.Printf(tableRowFormat, "firstPointer", firstPointer, &firstPointer, firstPointer)
	fmt.Printf(tableRowFormat, "*firstPointer", *firstPointer, &*firstPointer, *firstPointer)
	fmt.Println("|-----------------|----------|--------------|--------------|")
	fmt.Printf(tableRowFormat, "secondPointer", secondPointer, &secondPointer, secondPointer)
	fmt.Printf(tableRowFormat, "*secondPointer", *secondPointer, &*secondPointer, *secondPointer)
	fmt.Printf(tableRowFormat, "**secondPointer", **secondPointer, &**secondPointer, **secondPointer)
	fmt.Println("|-----------------|----------|--------------|--------------|")
	fmt.Printf("| %-15v | %-8T | %-12p | %-12p |\n", "thirdPointer", thirdPointer, &thirdPointer, thirdPointer)


	fmt.Println("\n------------------------------------------------------------------\n")

	anotherIntVar := 500

	fmt.Println("intVar        =", intVar)        // intVar        = 10
	fmt.Println("anotherIntVar =", anotherIntVar) // anotherIntVar = 500
	fmt.Println("*firstPointer =", *firstPointer) // *firstPointer = 10

	fmt.Println("-------------------")

	firstPointer = &anotherIntVar
	*firstPointer += 10

	fmt.Println("intVar        =", intVar)        // intVar        = 10
	fmt.Println("anotherIntVar =", anotherIntVar) // anotherIntVar = 510
	fmt.Println("*firstPointer =", *firstPointer) // *firstPointer = 510
}
