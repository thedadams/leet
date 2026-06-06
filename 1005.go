import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	var result, lastNeg int
	for _, n := range nums {
		if n < 0 {
			if k > 0 {
				result -= n
				k--
			} else {
				result += n
			}
			lastNeg = -n
		} else if n == 0 {
			k = 0
		} else {
			if k%2 == 1 {
				if lastNeg > 0 && lastNeg < n {
					result -= 2 * lastNeg
					result += n
				} else {
					result -= n
				}
				k = 0
			} else {
				result += n
			}
		}
	}

	if k%2 == 1 {
		// The only way to get here is if we have more indices to flip and everything was negative.
		// So take the bigest negative and flip it.
		result += 2 * nums[len(nums)-1]
	}

	return result
}
