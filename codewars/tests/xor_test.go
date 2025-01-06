package codewars_test

import (
	kata "codewars/kata"
	"testing"
)

func TestXor(t *testing.T) {
	// Define test cases
	testCases := []struct {
		a, b     bool
		expected bool
	}{
		{a: false, b: false, expected: false}, // Neither is true => false
		{a: false, b: true, expected: true},   // One is true => true
		{a: true, b: false, expected: true},   // One is true => true
		{a: true, b: true, expected: false},   // Both true => false
	}

	for _, tc := range testCases {
		result := kata.Xor(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Xor(%v, %v) = %v; expected %v",
				tc.a, tc.b, result, tc.expected)
		}
	}
}
