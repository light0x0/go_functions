package piscine

// import "github.com/01-edu/z01"
func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}

	return 1
}
