func addToArrayForm(num []int, k int) []int {
	var carry int
	idx := len(num) - 1
	for k > 0 || carry > 0 {
		digit := k % 10
		k /= 10

		if idx == -1 {
			num = append([]int{0}, num...)
			idx = 0
		}

		num[idx] += digit + carry
		carry = num[idx] / 10
		num[idx] %= 10
		idx--
	}

	return num
}
