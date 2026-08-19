package piscine

func NRune(s string, n int) rune {
	if n < 0 {
		return []rune(s)[len(s)+n]
	}
	return []rune(s)[n-1]
}
