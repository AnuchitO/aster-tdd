package main

import "fmt"

func FizzBuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}

	return fmt.Sprint(n)
}
