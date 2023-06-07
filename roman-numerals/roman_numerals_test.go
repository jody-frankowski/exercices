package roman_numerals

import "testing"

func TestUintToRoman(t *testing.T) {
	tests := []struct {
		input uint
		want  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{44, "XLIV"},
		{150, "CL"},
		{200, "CC"},
		{490, "CDXC"},
		{999, "CMXCIX"},
		{1984, "MCMLXXXIV"},
		{2001, "MMI"},
	}

	for _, test := range tests {
		got := UintToRoman(test.input)
		want := test.want
		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	}
}

func TestRomanToUint(t *testing.T) {
	tests := []struct {
		want uint
		input  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{44, "XLIV"},
		{150, "CL"},
		{200, "CC"},
		{490, "CDXC"},
		{999, "CMXCIX"},
		{1984, "MCMLXXXIV"},
		{2001, "MMI"},
	}

	for _, test := range tests {
		got := RomanToUint(test.input)
		want := test.want
		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	}
}
