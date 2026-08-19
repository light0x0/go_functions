package piscine

func Capitalize(s string) string {
	str := []rune(s)
	prev := false
	for index := 0; index < len(str); index++ {
		value := str[index]
		if (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') {
			if !prev {
				if value >= 'a' && value <= 'z' {
					str[index] = value - 32
				}
			}
			prev = true
		} else {
			prev = false
		}
	}
	return string(str)
}
