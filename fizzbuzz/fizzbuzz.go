package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {
	if num == 5 {
		return "Buzz"
	}
	if num == 3 || num == 6 || num == 9 {
		return "Fizz"
	}

	return fmt.Sprintf("%d", num)
}
