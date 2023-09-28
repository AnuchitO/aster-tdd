package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		input int
		want  string
		name  string
	}{
		{input: 1, want: "1", name: "should return 1 when input 1"},
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

func TestFizzBuzzShouldReturn2WhenInput2(t *testing.T) {
	input := 2

	got := FizzBuzz(input)

	want := "2"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput3(t *testing.T) {
	input := 3

	got := FizzBuzz(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
