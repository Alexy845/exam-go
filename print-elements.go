package exam

func PrintElements(tab []int, str []string) string {
	liste := ""
	for _, i := range tab {
		if i >= 0 && i < len(str) {
			liste += str[i]
		} else {
			liste += "!"
		}
		if liste != "" {
			liste += " "
		}
	}
	return liste
}
