package lesson_25_testing

// Simple function that explode string by space and filter empty strings
func explodeNonEmpty(data string) []string {
	var (
		result []string
		word string
		stringSymbol string
	)
	separator := " "

	for _, symbol := range data {
		stringSymbol = string(symbol)
		if stringSymbol != separator {
			word += stringSymbol
		} else {
			if len(word) > 0 {
				result = append(result, word)
				word = ""
			}
		}
	}

	if len(word) > 0 {
		result = append(result, word)
	}

	return result
}
