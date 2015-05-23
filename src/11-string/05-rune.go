package main

import "fmt"

func main() {
// START OMIT
String := "АБВ"
Rune := []rune(String)

SymbolAB := string(Rune[:2])
// END OMIT

	fmt.Printf("%-6s (%T): len: %v\nvalue: %v\n\n", "String", String, len(String), String)
	fmt.Printf("%-6s (%T): len: %v\nvalue: %v\n\n", "Rune", Rune, len(Rune), Rune)
	fmt.Printf("%-6s (%T): len: %v\nvalue: %v", "SymbolAB", SymbolAB, len(SymbolAB), SymbolAB)
}
