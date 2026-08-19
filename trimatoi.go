package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	first := false
	for _, value := range s {
		if !first {
			if value == '-' {
				sign = -1
			}
			if value == '+' {
				sign = 1
			}
		}
		if value >= '0' && value <= '9' {
			first = true
			num = num*10 + int(value-'0')
		}
	}

	return num * sign
}
