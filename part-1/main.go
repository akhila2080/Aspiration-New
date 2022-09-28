package main

import (
	"fmt"
	"unicode"
)

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
		if unicode.IsLetter(charAtPosition) || unicode.IsDigit(charAtPosition) {
			alphanumericCharCount += 1
			if alphanumericCharCount%3 == 0 {
				output[index] = unicode.ToUpper(charAtPosition)
			} else {
				output[index] = unicode.ToLower(charAtPosition)
			}
		} else {
			output[index] = charAtPosition
		}

	}
	return string(output)
}
