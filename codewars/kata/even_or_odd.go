package codewars

func EvenOrOdd(number int) string {
	return []string{"Even", "Odd"}[number&1]
}
