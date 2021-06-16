package roman

import (
	"testing"
)

func TestFromArabic(t *testing.T) {
	tests := []struct{
		input int
		expected string
	} {
		{-1, ""},
		{1, "I"},
		{10, "X"},
		{35, "XXXV"},
		{994, "CMXCIV"},
		{993, "CMXCIII"},
		{523, "DXXIII"},
		{325, "CCCXXV"},
		{918, "CMXVIII"},
		{990, "CMXC"},
		{859, "DCCCLIX"},
		{985, "CMLXXXV"},
		{629, "DCXXIX"},
		{855, "DCCCLV"},
		{172, "CLXXII"},
		{554, "DLIV"},
		{207, "CCVII"},
		{61, "LXI"},
		{72, "LXXII"},
		{615, "DCXV"},
		{986, "CMLXXXVI"},
		{362, "CCCLXII"},
		{756, "DCCLVI"},
		{892, "DCCCXCII"},
		{24, "XXIV"},
		{803, "DCCCIII"},
		{368, "CCCLXVIII"},
		{393, "CCCXCIII"},
		{82, "LXXXII"},
		{791, "DCCXCI"},
		{909, "CMIX"},
		{75, "LXXV"},
		{726, "DCCXXVI"},
		{898, "DCCCXCVIII"},
		{447, "CDXLVII"},
		{990, "CMXC"},
		{237, "CCXXXVII"},
		{482, "CDLXXXII"},
		{864, "DCCCLXIV"},
		{478, "CDLXXVIII"},
		{588, "DLXXXVIII"},
		{885, "DCCCLXXXV"},
		{96, "XCVI"},
		{369, "CCCLXIX"},
		{794, "DCCXCIV"},
		{212, "CCXII"},
		{237, "CCXXXVII"},
		{775, "DCCLXXV"},
		{542, "DXLII"},
		{854, "DCCCLIV"},
		{798, "DCCXCVIII"},
		{381, "CCCLXXXI"},
		{447, "CDXLVII"},
		{543, "DXLIII"},
		{515, "DXV"},
		{63, "LXIII"},
		{595, "DXCV"},
		{3999, "MMMCMXCIX"},
		{4000, ""},
	}

	for _, test := range tests {
		output := FromArabic(test.input)

		if output != test.expected {
			t.Errorf("Expected '%d' to be '%s' but got '%s'", test.input, test.expected, output)
		}
	}
}

func TestToArabic(t *testing.T) {
	tests := []struct{
		input string
		expected int
	} {
		{"I", 1},
		{"IIII", 4},
		{"IV", 4},
		{"DCCCLVI",856},
		{"CMXLI", 941},
		{"CXV", 115},
		{"CMXCIV", 994},
		{"", 0},
		{"MMMCMXCIX", 3999},
		{"CCCXXV", 325},
		{"CMXVIII", 918},
		{"CMXC", 990},
		{"DCCCLIX", 859},
		{"CMLXXXV", 985},
		{"DCXXIX", 629},
		{"DCCCLV", 855},
		{"CLXXII", 172},
		{"DLIV", 554},
		{"CCVII", 207},
		{"LXI", 61},
		{"LXXII", 72},
		{"DCXV", 615},
		{"CMLXXXVI", 986},
		{"CCCLXII", 362},
		{"DCCLVI", 756},
		{"DCCCXCII", 892},
		{"XXIV", 24},
		{"DCCCIII", 803},
		{"CCCLXVIII", 368},
		{"CCCXCIII", 393},
		{"LXXXII", 82},
		{"DCCXCI", 791},
		{"CMIX", 909},
		{"LXXV", 75},
		{"DCCXXVI", 726},
		{"DCCCXCVIII", 898},
		{"CDXLVII", 447},
		{"CMXC", 990},
		{"CCXXXVII", 237},
		{"CDLXXXII", 482},
		{"DCCCLXIV", 864},
		{"CDLXXVIII", 478},
		{"DLXXXVIII", 588},
		{"DCCCLXXXV", 885},
		{"XCVI", 96},
		{"CCCLXIX", 369},
		{"DCCXCIV", 794},
		{"CCXII", 212},
		{"CCXXXVII", 237},
		{"DCCLXXV", 775},
		{"DXLII", 542},
		{"DCCCLIV", 854},
		{"DCCXCVIII", 798},
		{"CCCLXXXI", 381},
		{"CDXLVII", 447},
		{"DXLIII", 543},
		{"DXV", 515},
		{"LXIII", 63},
		{"DXCV", 595},
	}

	for _, test := range tests {
		output := ToArabic(test.input)

		if output != test.expected {
			t.Errorf("Expected '%s' to be '%d' but got '%d'", test.input, test.expected, output)
		}
	}
}