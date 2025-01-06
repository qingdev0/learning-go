package codewars

func RockPaperScissors(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}

	wins := map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}

	if wins[p1] == p2 {
		return "Player 1 won!"
	}
	return "Player 2 won!"
}
