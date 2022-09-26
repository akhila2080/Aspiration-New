package main

import (
	"fmt"
)

// Interface StringMapper contains two functions. THis interface can be implemented to apply any
// transformation on the string
type StringMapper interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

/*
	This struct implements the interface to apply the transformation "capitalize every nth alphanumeric character on

any given string"
*/
type CapitalizeEveryThirdAlphanumericCharMapper struct {
	Entity  string
	Intervl int
	Counter int
	Output  []rune
}

func main() {
	str := NewCapitalizeEveryThirdAlphanumericCharMapper(3, "Aspiration.com")
	MapString(&str)
	fmt.Println(string(str.Output))
}

// NewEntity is used to build initial payload of an Input model.
func NewCapitalizeEveryThirdAlphanumericCharMapper(l int, str string) CapitalizeEveryThirdAlphanumericCharMapper {
	return CapitalizeEveryThirdAlphanumericCharMapper{
		Entity:  str,
		Intervl: l,
		Counter: 1,
		Output:  make([]rune, len(str)),
	}
}

// MapString is used to process the runes using for loop.
func MapString(input StringMapper) {
	for pos := range input.GetValueAsRuneSlice() {
		input.TransformRune(pos)
	}
}

// TransformRune finds the character at position parameter and applies following logic to it,
// 1. converts to upper case if it is third alphanumeric character
// 2. converts to lower case if it alphanumeric but not third character in a row
// 3. doesn't transform if not alphanumeric character
func (str *CapitalizeEveryThirdAlphanumericCharMapper) TransformRune(pos int) {
	if pos >= len(str.Entity) {
		return
	}
	charAtPosition := rune(str.Entity[pos])
	if !isAlphanumeric(charAtPosition) {
		str.Output[pos] = charAtPosition
		return
	}

	// if rune is at required interval, capitalize the rune.
	if str.Counter%str.Intervl == 0 {
		str.Output[pos] = toUpper(charAtPosition)
	} else {
		str.Output[pos] = toLower(charAtPosition)
	}

	// increase the counter value.
	str.Counter++
}

// GetValueAsRuneSlice will return the slice of rune values from input string.
func (str *CapitalizeEveryThirdAlphanumericCharMapper) GetValueAsRuneSlice() []rune {
	return []rune(str.Entity)
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
