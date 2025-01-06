package codewars_test

import (
	kata "codewars/kata"
	"testing"
)

func TestRockPaperScissors(t *testing.T) {
	tests := []struct {
		p1, p2 string
		want   string
	}{
		{"rock", "rock", "Draw!"},
		{"rock", "paper", "Player 2 won!"},
		{"rock", "scissors", "Player 1 won!"},
		{"paper", "rock", "Player 1 won!"},
		{"paper", "paper", "Draw!"},
		{"paper", "scissors", "Player 2 won!"},
		{"scissors", "rock", "Player 2 won!"},
		{"scissors", "paper", "Player 1 won!"},
		{"scissors", "scissors", "Draw!"},
	}

	for _, tt := range tests {
		t.Run(tt.p1+" vs "+tt.p2, func(t *testing.T) {
			got := kata.RockPaperScissors(tt.p1, tt.p2)
			if got != tt.want {
				t.Errorf(
					"RockPaperScissors(%q, %q) = %q; want %q",
					tt.p1, tt.p2, got, tt.want)
			}
		})
	}
}
