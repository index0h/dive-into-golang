package main

import "fmt"

func main() {
// START OMIT
String := `some here`

SSymbol := String[0]
SCode := string(SSymbol)
Some := String[:4]
Here := String[5:]

String = Some + " string " + Here
// END OMIT

	fmt.Printf("%-7s (%-6T): %v\n", "SSymbol", SSymbol, SSymbol)
	fmt.Printf("%-7s (%-6T): %v\n", "SCode", SCode, SCode)
	fmt.Printf("%-7s (%-6T): %v\n", "Some", Some, Some)
	fmt.Printf("%-7s (%-6T): %v\n\n", "Here", Here, Here)
	fmt.Printf("%s (%T):\nvalue: %v", "String", String, String)
}
