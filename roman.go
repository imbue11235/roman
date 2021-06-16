// Package roman provides functions for
// converting arabic to roman numerals and back
package roman

import (
	"strings"
)

// Largest value for roman numerals.
// "The largest number that can be represented in this notation is 3,999 (MMMCMXCIX)"
// (https://en.wikipedia.org/wiki/Roman_numerals)
const upperbound = 3999

type numeral struct {
	symbol string
	value int
}

var numerals = []numeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// FromArabic takes an arabic numeral and converts it to
// a roman numeral. If the arabic numeral is 0 or higher than
// set upperbound (3999), it will return an empty string.
func FromArabic(arabic int) (output string) {
	if arabic == 0 || arabic > upperbound {
		return
	}

	for _, numeral := range numerals {
		quotient := arabic/numeral.value

		output += strings.Repeat(numeral.symbol, quotient)
		arabic -= numeral.value * quotient

		if arabic <= 0 {
			break
		}
	}

	return
}

// ToArabic takes a roman numeral and converts it to
// an arabic numeral (int).
func ToArabic(roman string) (output int) {
	for _, numeral := range numerals {
		for strings.Index(roman, numeral.symbol) == 0 {
			output += numeral.value
			roman = roman[len(numeral.symbol):]
		}
	}

	return
}