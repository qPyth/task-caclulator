package number

import "strings"

func ArabicToRoman(num int) string {
	var builder strings.Builder

	romanMap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	ranks := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, rank := range ranks {
		for num >= rank {
			builder.WriteString(romanMap[rank])
			num -= rank
		}
	}

	return builder.String()
}

func RomanToArabic(s string) int {
	nums := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	var res int
	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			res += nums[s[i]]
		} else {
			if nums[s[i]] < nums[s[i+1]] {
				res -= nums[s[i]]
			} else {
				res += nums[s[i]]
			}
		}
	}
	return res
}
