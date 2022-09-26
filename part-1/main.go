package main

import "fmt"

func main() {
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))

}

// CapitalizeEveryThirdAlphanumericChar applies the following transformations in any given string,
// 1. converts every third alphanumeric character to uppercase
// 2. converts rest of the alphanumeric characters to lowercase
// 3. no operation on non alphanumeric characters
// Example,
//
//	Input:"Aspiration.com"
//	Output:"asPirAtiOn.cOm"
func CapitalizeEveryThirdAlphanumericChar(str string) string {
	output := make([]rune, len(str))
	alphanumericCharCount := 0
	for index, charAtPosition := range str {
		if isAlphanumeric(charAtPosition) {
			alphanumericCharCount += 1
			if alphanumericCharCount%3 == 0 {
				output[index] = toUpper(charAtPosition)
			} else {
				output[index] = toLower(charAtPosition)
			}
		} else {
			output[index] = charAtPosition
		}

	}
	return string(output)
}

// isAlphanumeric returns true if character belongs to the set of alphanumeric characters ([a-z,A-Z,0-9])
func isAlphanumeric(character rune) bool {
	return (character >= '0' && character <= '9') || (character >= 'A' && character <= 'Z') || (character >= 'a' && character <= 'z')
}

// toUpper converts any alphanumeric rune to its upper case form
func toUpper(character rune) rune {
	if character >= 'a' && character <= 'z' {
		return character - 'a' + 'A'
	}
	return character
}

// toLower converts any alphanumeric rune to its lower case form
func toLower(character rune) rune {
	if character >= 'A' && character <= 'Z' {
		return character - 'A' + 'a'
	}
	return character
}
