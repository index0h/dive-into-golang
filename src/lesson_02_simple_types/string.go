package lesson_02_simple_types

import (
	"fmt"
)

// \\			Обратный слеш ( \ )
// \OOO			8-битный символ Юникода, заданный тремя шестнадцатеричными цифрами
// \'			Апостроф ( ' )
// \"			Кавычки ( " )
// \a			ASCII-символ «сигнал» (BEL)
// \b			ASCII-символ «забой» (BS)
// \f			ASCII-символ «перевод формата» (FF)
// \n			ASCII-символ «перевод строки» (LF)
// \r			ASCII-символ «возврат каретки» (CR)
// \t			ASCII-символ «табуляция» (TAB)
// \uhhhh		16-битный символ Юникода, заданный 4-мя шестнадцатеричными цифрами
// \Uhhhhhhhh	32-битный символ Юникода, заданный 8-ю шестнадцатеричными цифрами
// \v			ASCII-символ «вертикальная табуляция» (VT)
// \xhh			8-битный символ Юникода, заданный двумя шестнадцатеричными цифрами
func DemonstrateString() {
	first := "First: line 1\nline 2"
	second := `Second: line 1\nline 2`
	third := "Многоязычная string"

	fmt.Println("\t------------------------------------------------------------------\n\n")

	fmt.Println("\tСтроки\n")

	fmt.Println("first  = ", first) // First: line 1
	//                               //  line 2
	fmt.Println("second = ", second) // Second: line 1\nline 2
	fmt.Println(len(third))          // 31, а не 19

	stringOperations()
}

func stringOperations() {
	first := "first"
	second := "second"
	space := " "
	ukrainian := "українська"
	runeUkrainian := []rune(ukrainian)

	fmt.Println("\t------------------------------------------------------------------\n\n")

	fmt.Println("\tСтроковые операции\n")

	fmt.Println("first + space + second  = ", first+space+second)                // "first second"
	first += space + second;	fmt.Println("first += space + second = ", first) // "first second"
	fmt.Println("first[1]                = ", first[1])                          // 105
	fmt.Println("string(first[1])        = ", string(first[1]))                  // "i"
	fmt.Println("first[3:]               = ", first[3:])                         // "st second"
	fmt.Println("first[:7]               = ", first[:7])                         // "first s"
	fmt.Println("first[2:7]              = ", first[2:7], "\n")                  // "rst s"

	fmt.Println("ukrainian[1]            = ", ukrainian[1])         // 131
	fmt.Println("ukrainian[3:]           = ", ukrainian[3:])        // "?раїнська"
	fmt.Println("ukrainian[:7]           = ", ukrainian[:7])        // "укр?"
	fmt.Println("ukrainian[2:7]          = ", ukrainian[2:7], "\n") // "кр?"

	fmt.Println("runeUkrainian[1]            = ", runeUkrainian[1])   // 1082
	fmt.Println("runeUkrainian[3:]           = ", runeUkrainian[3:])  // [1072 1111 1085 1089 1100 1082 1072]
	fmt.Println("runeUkrainian[:7]           = ", runeUkrainian[:7])  // [1091 1082 1088 1072 1111 1085 1089]
	fmt.Println("runeUkrainian[2:7]          = ", runeUkrainian[2:7]) // [1088 1072 1111 1085 1089]

	fmt.Println("\nstring(runeUkrainian[1])  = ", string(runeUkrainian[1]))         // "к"
	fmt.Println("string(runeUkrainian[3:])   = ", string(runeUkrainian[3:]))        // "аїнська"
	fmt.Println("string(runeUkrainian[:7])   = ", string(runeUkrainian[:7]))        // "українс"
	fmt.Println("string(runeUkrainian[2:7])  = ", string(runeUkrainian[2:7]), "\n") // "раїнс"
}
