// test codewars/kata/main.go

package codewars_test

import (
	kata "codewars/kata"
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{2, "Even"},
		{3, "Odd"},
	}

	for _, test := range tests {
		got := kata.EvenOrOdd(test.input)
		if got != test.expected {
			t.Errorf("EvenOrOdd(%d) = %s; want %s", test.input, got, test.expected)
		}
	}
}
