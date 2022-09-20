package piscine

func ConcatParams(args []string) string {
	str := ""
	for i := 0; i <= len(args)-1; i++ {
		str += args[i]
		if len(args)-1 > i {
			str += "\n"
		}

	}
	return str
}
