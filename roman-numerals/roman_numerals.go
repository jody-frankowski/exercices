package roman_numerals

import "strings"

type symbol struct {
	value  uint
	symbol string
}
var symbols = []symbol{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func UintToRoman(arabic uint) string {
	roman := ""

	for _, symbol := range symbols {
		for arabic >= symbol.value {
			roman += symbol.symbol
			arabic -= symbol.value
		}
	}
	return roman
}

func RomanToUint(roman string) uint {
	var arabic uint = 0

	for len(roman) > 0 {
		for _, symbol := range symbols {
			if strings.HasPrefix(roman, symbol.symbol) {
				arabic += symbol.value
				roman = strings.TrimPrefix(roman, symbol.symbol)
				break
			}
		}
	}
	return arabic
}
