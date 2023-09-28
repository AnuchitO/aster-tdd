package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		input int
		want  string
		name  string
	}{
		{input: 1, want: "1", name: "should return 1 when input 1"},
		{input: 2, want: "2", name: "should return 2 when input 2"},
		{input: 3, want: "Fizz", name: "should return Fizz when input 3"},
		{input: 4, want: "4", name: "should return 4 when input 4"},
		{input: 5, want: "Buzz", name: "should return Buzz when input 5"},
		{input: 6, want: "Fizz", name: "should return Fizz when input 6"},
		{input: 7, want: "7", name: "should return 7 when input 7"},
		{input: 8, want: "8", name: "should return 8 when input 8"},
		{input: 9, want: "Fizz", name: "should return Fizz when input 9"},
		{input: 10, want: "Buzz", name: "should return Buzz when input 10"},
		{input: 11, want: "11", name: "should return 11 when input 11"},
		{input: 12, want: "Fizz", name: "should return Fizz when input 12"},
		{input: 13, want: "13", name: "should return 13 when input 13"},
		{input: 14, want: "14", name: "should return 14 when input 14"},
		{input: 15, want: "FizzBuzz", name: "should return FizzBuzz when input 15"},
		{input: 30, want: "FizzBuzz", name: "should return FizzBuzz when input 30"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FizzBuzz(c.input)

			if got != c.want {
				t.Errorf("got %q but want %q", got, c.want)
			}
		})
	}
}
