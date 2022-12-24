package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	input := 1
	want := "1"

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(1) = %q, want %q", got, want)
	}
}
func TestFizzBuzz2ShouldGet2(t *testing.T) {
	input := 2
	want := "2"

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(1) = %q, want %q", got, want)
	}
}

func TestFizzBuzz3ShouldGetFizz(t *testing.T) {
	input := 3
	want := "Fizz"

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(1) = %q, want %q", got, want)
	}
}

func TestFizzBuzz4ShouldGet4(t *testing.T) {
	input := 4
	want := "4"

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(1) = %q, want %q", got, want)
	}
}
