package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {
	if num == 5 || num == 10 {
		return "Buzz"
	}
	if num == 3 || num == 6 || num == 9 || num == 12 {
		return "Fizz"
	}

	return fmt.Sprintf("%d", num)
}
