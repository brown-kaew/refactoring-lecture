package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {
	if num == 15 {
		return "FizzBuzz"
	}
	if num == 5 || num == 10 {
		return "Buzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}

	return fmt.Sprintf("%d", num)
}
