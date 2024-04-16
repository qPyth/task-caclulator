package number

import "testing"

var tests = []struct {
	arabic int
	roman  string
}{
	{1, "I"},
	{4, "IV"},
	{9, "IX"},
	{58, "LVIII"},
	{1994, "MCMXCIV"},
	{3, "III"},
	{7, "VII"},
	{11, "XI"},
	{44, "XLIV"},
	{99, "XCIX"},
	{500, "D"},
	{1000, "M"},
	{3999, "MMMCMXCIX"},
	{0, ""},
}

func TestArabicToRoman(t *testing.T) {

	for _, test := range tests {
		res := ArabicToRoman(test.arabic)

		if res != test.roman {
			t.Errorf("arabicToRoman(%d) = %s; want %s", test.arabic, res, test.roman)
		}
	}
}

func TestRomanToArabic(t *testing.T) {
	for _, test := range tests {
		res := RomanToArabic(test.roman)

		if res != test.arabic {
			t.Errorf("arabicToRoman(%s) = %d; want %d", test.roman, res, test.arabic)
		}
	}
}
