package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{input: 1, want: "1"},
		{input: 2, want: "2"},
		{input: 3, want: "Fizz"},
		{input: 4, want: "4"},
		{input: 5, want: "Buzz"},
		{input: 6, want: "Fizz"},
		{input: 7, want: "7"},
		{input: 8, want: "8"},
		{input: 9, want: "Fizz"},
		{input: 10, want: "Buzz"},
		{input: 11, want: "11"},
		{input: 12, want: "Fizz"},
		{input: 13, want: "13"},
		{input: 14, want: "14"},
		{input: 15, want: "FizzBuzz"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d should get %s", tt.input, tt.want), func(t *testing.T) {
			got := FizzBuzz(tt.input)

			if got != tt.want {
				t.Errorf("got %q, but want %q", got, tt.want)
			}
		})
	}
}
