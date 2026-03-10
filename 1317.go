func getNoZeroIntegers(n int) []int {
	var a, b = n - 1, 1

	for {
		if !hasZeroDigit(a) && !hasZeroDigit(b) {
			return []int{a, b}
		}

        a--
        b++
	}
}

func hasZeroDigit(n int) bool {
    for n > 0 {
        if n % 10 == 0 {
            return true
        }

        n /= 10
    }

    return false
}