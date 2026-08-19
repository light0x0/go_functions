package piscine

func NRune(s string, n int) rune {
	if len(s) == 0 {
		return 0
	}
	if n < 0 {
		return []rune(s)[len(s)+n]
	}
	return []rune(s)[n-1]
}
