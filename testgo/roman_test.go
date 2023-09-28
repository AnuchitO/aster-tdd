package main

import "testing"

func TestRomanToArabic(t *testing.T) {
	cases := []struct {
		input string
		want  int
		name  string
	}{
		{input: "I", want: 1, name: "should return 1 when input I"},
		{input: "II", want: 2, name: "should return 2 when input II"},
		{input: "III", want: 3, name: "should return 3 when input III"},
		{input: "IV", want: 4, name: "should return 4 when input IV"},
		{input: "V", want: 5, name: "should return 5 when input V"},
		{input: "VI", want: 6, name: "should return 6 when input VI"},
		{input: "VII", want: 7, name: "should return 7 when input VII"},
		{input: "VIII", want: 8, name: "should return 8 when input VIII"},
		{input: "IX", want: 9, name: "should return 9 when input IX"},
		{input: "X", want: 10, name: "should return 10 when input X"},
		{input: "XXIX", want: 29, name: "should return 29 when input XXIX"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := RomanToArabic(c.input)

			if got != c.want {
				t.Errorf("got %d but want %d", got, c.want)
			}
		})
	}

}
