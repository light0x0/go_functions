package piscine

func ToLower(s string) string {
	splited := []rune(s)
	for index, value := range s {
		if value >= 'A' && value <= 'Z' {
			splited[index] = value + 32
		}
	}
	return string(splited)
}
