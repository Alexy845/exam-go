package exam

func Decipher(s string) string {
	str := ""
	for _, i := range s {
		if i > 96 && i < 123 {
			if i-16 < 97 {
				i += 26
			}
			str += string(i - 16)

		} else if i > 66 && i < 91 {
			if i-16 < 65 {
				i += 26
			}
			str += string(i - 16)
		} else {
			str += string(i)
		}
	}
	return str
}
