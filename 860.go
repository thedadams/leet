func lemonadeChange(bills []int) bool {
	change := make(map[int]int, 3)
	for _, b := range bills {
		change[b]++
		b -= 5
		if b >= 10 && change[10] > 0 {
			b -= 10
			change[10]--
		}

		change[5] -= b / 5
		if change[5] < 0 {
			return false
		}
	}

	return true
}
