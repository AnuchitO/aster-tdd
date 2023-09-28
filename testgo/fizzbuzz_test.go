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
