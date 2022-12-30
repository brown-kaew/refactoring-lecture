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
