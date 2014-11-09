package main

import (
	"fmt"
	"strings"
)

type word string

// Check if there is not word
func (this *word) isEmpty() bool {
	return (*this == "")
}

// Wrap word by symbols
func (this *word) wrap(wrapSymbols string) *word {
	*this = word(wrapSymbols + string(*this) + wrapSymbols)

	return this
}

// Remove space symbols form begin and and of word
func (this *word) trim() *word {
	*this = word(strings.TrimSpace(string(*this)))

	return this
}

func main() {
	testWord := word("   test   ")

	testWord.trim().wrap("+")

	fmt.Println(`word.isEmpty():        `, testWord.isEmpty())
	fmt.Println(`word.trim().wrap("+"): `, testWord)
}
