package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		digit := n % 10
		n /= 10
		z01.PrintRune(rune('0' + digit))
	}
}
