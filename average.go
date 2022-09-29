package exam

func Average(tab []int) int {
	diviseur := 0
	moyenne := 0
	for _, i := range tab {
		if i == -1 {
			continue
		} else {
			diviseur += 1
			moyenne += i
		}
	}
	return moyenne / diviseur
}
