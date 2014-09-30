package main

import (
	"fmt"
	"math"
)

func main() {
	var defaultInt int64
	intVar := 10

	fmt.Println("2 + 3        = ", 2 + 3)            // 5
	fmt.Println("2 - 3        = ", 2 - 3)            // -1
	fmt.Println("2 * 3        = ", 2 * 3)            // 6
	fmt.Println("int(2 / 3)   = ", int(2 / 3))       // 0
	fmt.Println("defaultInt   = ", defaultInt, "\n") // 0

	intVar--;		fmt.Println("intVar--     = ", intVar) // 9
	intVar++;		fmt.Println("intVar++     = ", intVar) // 10
	intVar += 10;	fmt.Println("intVar += 10 = ", intVar) // 20
	intVar -= 5;	fmt.Println("intVar -= 5  = ", intVar) // 15
	intVar *= 3;	fmt.Println("intVar *= 3  = ", intVar) // 45
	intVar /= 7;	fmt.Println("intVar /= 7  = ", intVar) // 6

	integerSize()
	bitOperations()
	changeIntegerType()
}

func integerSize() {
	fmt.Println("\n------------------------------------------------------------------\n")


	fmt.Println("\tПределы значений целых int*\n")

	fmt.Printf("int8:   [%v, %v]\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16:  [%v, %v]\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32:  [%v, %v]\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64:  [%v, %v]\n", math.MinInt64, math.MaxInt64)

	fmt.Println("\n\tПределы значений целых uint*\n")

	fmt.Printf("uint8:  [0, %v]\n", math.MaxUint8)
	fmt.Printf("uint16: [0, %v]\n", math.MaxUint16)
	fmt.Printf("uint32: [0, %v]\n", math.MaxUint32)
	fmt.Printf("uint64: [0, %v]\n", ^uint64(0))

	fmt.Println("\n\tСинонимы целых типов\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, или int64, в зависимости от ОС")
	fmt.Println("uint    - uint32, или uint64, в зависимости от ОС")
	fmt.Println("uintptr - Беззнаковое целое, пригодное для хранения значения указателя")
}

func bitOperations() {
	var x, y, z uint8

	x = 9
	y = 28
	z = x

	fmt.Println("\t------------------------------------------------------------------\n\n")


	fmt.Println("\tБитовые операции\n")

	fmt.Printf("^x      = ^(%d)      = ^(%.8b)            = %.8b = %d\n", x, x, ^x, ^x)
	fmt.Printf("x << 2  = (%d << 2)  = (%.8b << 2)        = %.8b = %d\n", x, x, x << 2, x << 2)
	fmt.Printf("x >> 2  = (%d >> 2)  = (%.8b >> 2)        = %.8b = %d\n", x, x, x >> 2, x >> 2)
	fmt.Printf("x & y   = (%d & %d)  = (%.8b & %.8b)  = %.8b = %d\n", x, y, x, y, x & y, x & y)
	fmt.Printf("x | y   = (%d | %d)  = (%.8b | %.8b)  = %.8b = %d\n", x, y, x, y, x | y, x | y)
	fmt.Printf("x ^ y   = (%d ^ %d)  = (%.8b ^ %.8b)  = %.8b = %d\n", x, y, x, y, x ^ y, x ^ y)
	fmt.Printf("x &^ y  = (%d &^ %d) = (%.8b &^ %.8b) = %.8b = %d\n", x, y, x, y, x &^ y, x &^ y)
	fmt.Printf("x %% y   = (%d %% %d)  = (%.8b %% %.8b)  = %.8b = %d\n", x, y, x, y, x % y, x % y)

	fmt.Println("\n\tБитовые операции с присваиванием\n")

	x = z; x &= y;	fmt.Printf("x &= y   = (%d &= %d)  = (%.8b &= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z; x |= y;	fmt.Printf("x |= y   = (%d |= %d)  = (%.8b |= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z; x ^= y;	fmt.Printf("x ^= y   = (%d ^= %d)  = (%.8b ^= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	x = z; x &^= y;	fmt.Printf("x &^= y  = (%d &^= %d) = (%.8b &^= %.8b) = %.8b = %d\n", z, y, z, y, x, x)
	x = z; x %= y;	fmt.Printf("x %%= y   = (%d %%= %d)  = (%.8b %%= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
}

func changeIntegerType() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("\t------------------------------------------------------------------\n\n")


	fmt.Println("\tПриведение типов\n")

	fmt.Printf("variable8         = %-5v = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5v = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5v = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5v = %.16b\n", uint8(variable16), uint8(variable16))
}
