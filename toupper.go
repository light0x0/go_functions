package piscine

func ToUpper(s string) string {
	splited := []rune(s)
	for index, value := range splited {
		if value >= 'a' && value <= 'z' {
			splited[index] = value - 32
		}
	}
	return string(splited)
}
