package piscine

func BasicJoin(elems []string) string {
	str := ""
	for index := range elems {
		str += elems[index]
	}
	return str
}
