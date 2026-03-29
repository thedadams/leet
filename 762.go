var primes map[int]struct{} = map[int]struct{}{
	2:  struct{}{},
	3:  struct{}{},
	5:  struct{}{},
	7:  struct{}{},
	11: struct{}{},
	13: struct{}{},
	17: struct{}{},
	19: struct{}{},
	23: struct{}{},
	29: struct{}{},
	31: struct{}{},
	37: struct{}{},
}

func countPrimeSetBits(left int, right int) int {
	var result, idx int
	bits, count := bitsArr(left)
	_, ok := primes[count]
	if ok {
		result++
	}

	for range right - left {
		idx = 0
		for {
			if bits[idx] == 1 {
				bits[idx] = 0
				count--
			} else {
				bits[idx] = 1
				count++
				break
			}
			idx++
		}

		if _, ok = primes[count]; ok {
			result++
		}
	}

	return result
}

func bitsArr(num int) ([20]int, int) {
	var (
		bits     [20]int
		i, count int
	)
	for num > 0 {
		bits[i] = num % 2
		if bits[i] == 1 {
			count++
		}
		num >>= 1
		i++
	}
	return bits, count
}
