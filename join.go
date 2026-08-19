package piscine

func Join(strs []string, sep string) string {
	str := ""
	for index := range strs {
		str += strs[index] + sep
	}
	return str
}
