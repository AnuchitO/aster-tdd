package main

var numerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
}

func RomanToArabic(r string) int {
	arabic := 0
	prev := 0
	for _, c := range r {
		v := numerals[c]

		if v > prev {
			arabic = arabic + (v - 2*prev)
		} else {
			arabic = arabic + v
		}
		prev = v
	}

	return arabic
}
